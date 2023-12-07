#!/bin/bash
KEY_PAIR=tutorial-cluster
    ecs-cli up \
      --keypair $KEY_PAIR  \
      --capability-iam \
      --size 2 \
      --instance-type t3.medium \
      --tags project=converter-cluster,owner=denysf \
      --cluster-config converter \
      --ecs-profile converter