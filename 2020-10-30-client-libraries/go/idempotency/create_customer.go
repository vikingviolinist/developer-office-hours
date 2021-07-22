// For more information about this demo, please see:
//   https://www.youtube.com/watch?v=UA7c4Y524HY

package main

import (
  "github.com/stripe/stripe-go/v72"
  "github.com/stripe/stripe-go/v72/customer"
)

func main() {
  stripe.Key = "sk_test_xxx"

  params := &stripe.CustomerParams{
    Email: stripe.String("foo@bar.com"),
  }
  params.SetIdempotencyKey("a-long-random-string-20201014155429")

  customer.New(params)
}
