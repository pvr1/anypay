/*
 * AnyPay
 *
 * This the AnyPay service targeted towards, parents with children doing payments and russian oligarks
 *
 * API version: 1.0.0
 * Contact: per.von.rosen@swedbank.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

/*
import (
	"net/http"

	"github.com/gin-gonic/gin"
)
*/

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"

	"context"
	"flag"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	kafkaUtils "github.com/pvr1/anypay/kafka"

	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

var (
	// kafka
	kafkaBrokerURL     string
	kafkaVerbose       bool
	kafkaTopicIn       string
	kafkaTopicOut      string
	kafkaConsumerGroup string
	kafkaClientID      string
)

// AddFxorder - Add a new Foreig Exchange order
func AddFxorder(c *gin.Context) {
	//c.String(http.StatusOK, "FX Order place on market place\n")
	/*
		var byteMsg []byte
		byteMsg = make([]byte, 1024)
		// Read the Body content
		if c.Request.Body != nil {
			byteMsg, _ = ioutil.ReadAll(c.Request.Body)
		}

		// Restore the io.ReadCloser to its original state
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(byteMsg))

		log.Debug().Msg(string(byteMsg))

		// Continue to use the Body, like Binding it to a struct:
		var fxorder Fxorder
		c.Bind(&fxorder)
		//order := new(models.GeaOrder)
		//error := c.Bind(order)
	*/
	/*
		Other example
		if c.Request.Method == "POST" {
			var u User
			c.BindJSON(&u)
			c.JSON(http.StatusOK, gin.H{
				"user": u.Username,
				"pass": u.Password,
			})
		}
	*/

	var byteMsg []byte
	byteMsg = make([]byte, 1024)
	// Read the Body content
	if c.Request.Body != nil {
		byteMsg, _ = ioutil.ReadAll(c.Request.Body)
	}

	// Restore the io.ReadCloser to its original state
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(byteMsg))

	//log.Debug().Msg(string(byteMsg))

	var fxorder Fxorder
	if err := c.ShouldBindJSON(&fxorder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if kafkaBrokerURL == "" {
		//flag.StringVar(&kafkaBrokerURL, "kafka-brokers", "localhost:19092,localhost:29092,localhost:39092", "Kafka brokers in comma separated value")
		flag.StringVar(&kafkaBrokerURL, "kafka-brokers", "localhost:9092", "Kafka brokers in comma separated value")
		flag.BoolVar(&kafkaVerbose, "kafka-verbose", true, "Kafka verbose logging")
		flag.StringVar(&kafkaTopicIn, "kafka-topicIn", "foo2", "Kafka topic. Only one topic per worker.")
		flag.StringVar(&kafkaTopicOut, "kafka-topicOut", "foo2", "Kafka topic. Only one topic per worker.")
		flag.StringVar(&kafkaConsumerGroup, "kafka-consumer-group", "consumer-group", "Kafka consumer group")
		flag.StringVar(&kafkaClientID, "kafka-client-id", "my-client-id", "Kafka client id")
	}

	flag.Parse()

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	brokers := strings.Split(kafkaBrokerURL, ",")

	// make a new reader that consumes from topic-A
	config := kafka.ReaderConfig{
		Brokers:         brokers,
		GroupID:         kafkaClientID,
		Topic:           kafkaTopicIn,
		MinBytes:        10e3,            // 10KB
		MaxBytes:        10e6,            // 10MB
		MaxWait:         1 * time.Second, // Maximum amount of time to wait for new data to come when fetching batches of messages from kafka.
		ReadLagInterval: -1,
	}

	reader := kafka.NewReader(config)

	// connect to kafka
	kafkaProducer, err := kafkaUtils.Configure(strings.Split(kafkaBrokerURL, ","), kafkaClientID, kafkaTopicOut)
	if err != nil {
		log.Error().Str("error", err.Error()).Msg("unable to configure kafkaProducer")
		return
	}

	defer kafkaProducer.Close()
	defer reader.Close()

	log.Debug().Msg("Conected to kafka")

	m, err := reader.ReadMessage(context.Background())

	log.Debug().Msg("Read request message body")
	if err != nil {
		log.Error().Msgf("error while receiving message: %s", err.Error())
	}

	log.Debug().Msg("Inside for loop...")
	value := m.Value

	log.Debug().Msg("Got a Message")
	//		if m.CompressionCodec == snappy.NewCompressionCodec() {
	//			_, err = snappy.NewCompressionCodec().Decode(value, m.Value)
	//		}

	var ctx = context.Background()
	// log.Debug().Msg("************")
	// log.Debug().Msg("ID: " + string(fxorder.ID))
	// log.Debug().Msg("FX: " + fxorder.FX)
	// log.Debug().Msg("Amount: " + string(fxorder.Quantity))
	// log.Debug().Msg("SettlementDate: " + fxorder.SettlementDate.String())
	// log.Debug().Msg("OrderDate: " + fxorder.OrderDate.String())
	// log.Debug().Msg("************")
	err = kafkaUtils.Push(ctx, []byte(fxorder.FX), []byte(string(fxorder.Quantity)+"/"+fxorder.FX)) //fxorder skall parse:as
	if err != nil {
		log.Error().Msg("Kafka write to topic Out failed")
		c.String(http.StatusOK, "Kafka write to topic Out failed\n")
	} else {
		c.String(http.StatusOK, "Kafka write to topic Out worked\n")
	}

	log.Debug().Msg("Have pushed to kafka")
	if err != nil {
		log.Error().Msgf("error while receiving message: %s", err.Error())
	}
	log.Debug().Msg("Printing response from kafka")
	fmt.Printf("message at topic/partition/offset %v/%v/%v: %s\n", m.Topic, m.Partition, m.Offset, string(value))

	log.Debug().Msgf("Closing Down")

	//c.String(http.StatusOK, "FX Order place on market place\n")
}

// DeleteFxorder - Deletes a payment-instruction not settled yet
func DeleteFxorder(c *gin.Context) {
	c.String(http.StatusOK, "FX Order deleted. You were so close...\n")
}

// GetFxorders - Get all non-matched FX orders that the login has privileges to get
func GetFxorders(c *gin.Context) {
	c.String(http.StatusOK, "Get FX Order list. Offer:100MM EUR/SEK @ 10.53 , BID:100MM EUR/SEK @ 10.51 . \n")
}

// GetFxorder - Get a specific foreign exchange order (FX order) with details
func GetFxorder(c *gin.Context) {
	c.String(http.StatusOK, "Get specific FX Order status. You are best offer: 100MM EUR/SEK @ 10.53 . \n")
}

// UpdateFxorder - Update a existing  FX order not matched yet, on the market place
func UpdateFxorder(c *gin.Context) {
	c.String(http.StatusOK, "FX Order updated.")
}
