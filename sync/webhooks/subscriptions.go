package webhooks

import (
	"context"
	"encoding/json"
	"fmt"

	shopify "github.com/r0busta/go-shopify-graphql-model/graph/model"
	"github.com/r0busta/graphql"
	log "github.com/sirupsen/logrus"
	"gopkg.in/guregu/null.v4"
)

var shopifyTopics = map[Topic]shopify.WebhookSubscriptionTopic{
	TopicProductsCreate: shopify.WebhookSubscriptionTopicProductsCreate,
	TopicProductsUpdate: shopify.WebhookSubscriptionTopicProductsUpdate,
	TopicProductsDelete: shopify.WebhookSubscriptionTopicProductsDelete,
}

const LIST_WEBHOOKS = `
	query webhookSubscriptions($first: Int, $callbackUrl: URL, $format: WebhookSubscriptionFormat, $topics: [WebhookSubscriptionTopic!]) {
		webhookSubscriptions(first: $first, callbackUrl: $callbackUrl, format: $format, topics: $topics) {
			edges {
				node {
					id
					endpoint {
						... on WebhookHttpEndpoint {
						callbackUrl
						}
					}
					format
					topic
				}
			}
		}
	}
`

type Subscriptions struct {
	shopClient      *graphql.Client
	callbackBaseURL string
}

func (s *Subscriptions) ProvisionSubscriptions(topics []Topic) error {
	if s.callbackBaseURL == "" {
		return fmt.Errorf("hook proxy URL is not set")
	}

	format := shopify.WebhookSubscriptionFormatJSON
	for _, topic := range topics {
		shopifyTopic := shopifyTopics[topic]
		path := topicHandlerPaths[topic]
		callbackURL := fmt.Sprintf("%s%s%s", s.callbackBaseURL, routeGroupPrefix, path)
		exists, err := s.isSubscribed(callbackURL, shopifyTopic, format)
		if err != nil {
			return fmt.Errorf("error: %s", err)
		}
		if exists {
			continue
		}

		err = s.subscribe(callbackURL, shopifyTopic, format)
		if err != nil {
			return fmt.Errorf("error: %s", err)
		}
		log.Debugf("Subscribed to %s with the callback URL at %s", shopifyTopic, callbackURL)
	}
	return nil
}

func (s *Subscriptions) UnsubscribeFromAll() error {
	if s.callbackBaseURL == "" {
		return fmt.Errorf("hook proxy URL is not set")
	}

	subscriptions, err := s.listAllSubscriptions(s.callbackBaseURL)
	if err != nil {
		return fmt.Errorf("error listing webhook subscriptions: %s", err)
	}

	log.Tracef("Removing %d webhook subscriptions", len(subscriptions.Edges))
	for _, sub := range subscriptions.Edges {
		err := s.unsubscribe(sub.Node.ID.String)
		if err != nil {
			return fmt.Errorf("error removing webhook subscription: %s", err)
		}
	}

	return nil
}

func (s *Subscriptions) listAllSubscriptions(callbackBaseURL string) (*shopify.WebhookSubscriptionConnection, error) {
	vars := map[string]interface{}{
		"first": 50,
	}
	var out struct {
		WebhookSubscriptions *shopify.WebhookSubscriptionConnection `json:"webhookSubscriptions"`
	}
	err := s.shopClient.QueryString(context.Background(), LIST_WEBHOOKS, vars, &out)
	if err != nil {
		return nil, fmt.Errorf("error getting callbackUrl=%s webhook subscriptions: %s", callbackBaseURL, err)
	}

	return out.WebhookSubscriptions, nil
}

func (s *Subscriptions) isSubscribed(callbackUrl string, topic shopify.WebhookSubscriptionTopic, format shopify.WebhookSubscriptionFormat) (bool, error) {
	vars := map[string]interface{}{
		"first":       50,
		"topics":      []shopify.WebhookSubscriptionTopic{topic},
		"format":      format,
		"callbackUrl": callbackUrl,
	}
	var out struct {
		WebhookSubscriptions *shopify.WebhookSubscriptionConnection `json:"webhookSubscriptions"`
	}
	err := s.shopClient.QueryString(context.Background(), LIST_WEBHOOKS, vars, &out)
	if err != nil {
		return false, fmt.Errorf("error getting %s webhook subscriptions: %s", topic, err)
	}

	return len(out.WebhookSubscriptions.Edges) != 0, nil
}

func (s *Subscriptions) subscribe(callbackUrl string, topic shopify.WebhookSubscriptionTopic, format shopify.WebhookSubscriptionFormat) error {
	u := null.StringFrom(callbackUrl)
	webhookInput := shopify.WebhookSubscriptionInput{
		CallbackURL: &u,
		Format:      &format,
	}

	vars := map[string]interface{}{
		"topic": topic,
		"input": webhookInput,
	}

	var webhookSubscriptionCreate struct {
		WebhookSubscriptionCreatePayload shopify.WebhookSubscriptionCreatePayload `graphql:"webhookSubscriptionCreate(topic: $topic, webhookSubscription: $input)" json:"webhookSubscriptionCreate"`
	}
	err := s.shopClient.Mutate(context.Background(), &webhookSubscriptionCreate, vars)
	if err != nil {
		return fmt.Errorf("error creating %s webhook: %s", topic, err)
	}

	if len(webhookSubscriptionCreate.WebhookSubscriptionCreatePayload.UserErrors) > 0 {
		errs, _ := json.MarshalIndent(webhookSubscriptionCreate.WebhookSubscriptionCreatePayload.UserErrors, "", "    ")
		return fmt.Errorf("error creating %s webhook: %s", topic, errs)
	}

	return nil
}

func (s *Subscriptions) unsubscribe(id string) error {
	vars := map[string]interface{}{
		"id": id,
	}

	var webhookSubscriptionDelete struct {
		WebhookSubscriptionDeletePayload shopify.WebhookSubscriptionDeletePayload `graphql:"webhookSubscriptionDelete(id: $id)" json:"webhookSubscriptionDelete"`
	}
	err := s.shopClient.Mutate(context.Background(), &webhookSubscriptionDelete, vars)
	if err != nil {
		return fmt.Errorf("error creating %s webhook: %s", id, err)
	}

	if len(webhookSubscriptionDelete.WebhookSubscriptionDeletePayload.UserErrors) > 0 {
		errs, _ := json.MarshalIndent(webhookSubscriptionDelete.WebhookSubscriptionDeletePayload.UserErrors, "", "    ")
		return fmt.Errorf("error creating %s webhook: %s", id, errs)
	}

	return nil
}
