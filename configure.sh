#!/bin/bash
set -e
PROFILE_NAME=converter
CLUSTER_NAME=converter-cluster
REGION=us-east-1
LAUNCH_TYPE=EC2
ecs-cli configure profile --profile-name "$PROFILE_NAME" --access-key "$AWS_ACCESS_KEY_ID" --secret-key "$AWS_SECRET_ACCESS_KEY"
ecs-cli configure --cluster "$CLUSTER_NAME" --default-launch-type "$LAUNCH_TYPE" --region "$REGION" --config-name "$PROFILE_NAME"

# eyJwYXlsb2FkIjoibXJOVnNWWk9HWEJRZS9Bc21nV1dyQ2w2NFBrZDRJZ2pyREZyckgvVmJEN1RJL2hHa1Jac0NaVk1XM3FWZHNxR1BseFFRY05VeXY3Ly94a0JKMURlUmxmWDFVV0k1TkQ4ZWNPOCtONmZLMDFBZjFlRXlpMXNkNmlkWGZ0MU9XMXRCM3l1T21hSWdBcWk3NU9kK25xL3JXYzV3Q0pPanQrc3A5MytYM2hjQzJQQS9IcTUvME5NN2s5ZS91eU45blRIYjMrVzdYTlRhellkcTJzaVY2dXlMRGRPTkFyYWpGSzVneEQ2bjBkdFFmRDlsUyt4dHBGYVFNTUlFZFhvUDRVSm5TdExvSk9tdld5Nkl2dGh0bW16OXF3WUFlNENjVDlEN3NQRXNqY0dUUUdzbWdCUXZ6alhEZUVGWDUzbjd4U2tEbVJJdndCTlZYU3IrUlI4aWFvZ1FPckRuVHVCbXdMQTRwVXh6eHh5MEtYbWt1UlgyYnVrZnhFNmV4MC9jaWpvNDQ3ZFNSOFp5Nm5Ia2ZKbzRaRHZ6NDVKSzdtN01lSG4yWUdQM1hlcFh0MHo0M1pWMTBiMHlJaUlDQUFvS1BHQ1AxQm8zQ0k5Q2tuZHhtY3dKUkR0aDNxM3lSb2VuM0hiNVlzckF1ckkrUko4dHQ1aFdGL2xjaU1QVUxNQmtudDJ6ZER1OVhSNXdnK1QyRmx4cTRtaVpnTWU0bU01aDhmSmcxSWVEZEY5U0ZyMWFFVmtnSko0TDlWT1JuQ292Y2VxQTFQNlZwS0c4aXQ4T1l6aVk2NFRjYkx6ZmlOYmxUU1hCOGFFQlFaYWMvUjl4V3hERXd3Ym5ORUVpNTd5bUkralJ4cUpYcENFbFFWc3UwTjRBcVJySzFhcWdMemFzNGFHT2w4VHIwQmJ5dXd0ZGVLUVlwL2Zpb1ZOelNhZ1NrYTNQakRmRlVQeXI2UGRTbnBxTXlWMEZZNldweFBsSEpRek5XTmVmWmpQRXJyNDhrbDRNVU1xYnl1YkVhRkJmaTdLc2hDMmVDR3JWZHNFdXBub05TYy9vTTVDb3AwN3hjQ1VtSzcxc0drdEtSc1ZVQm1XR0I5WUE0SjB2cm5nTnhPSm1xZGgrOXhJblZYSkRPLzdVVXg0K2luZjBTM0I1Snltdi9ENE0vaGV0YURVNnptYVBEeXQyQWd4ZU9sZG93WWZsOXg0b1lYcHpnK0dJU1l6Z21JblM1dmNxRnFhbTMwUkFWaWg0MXJPUWtYU0ZFMFBDSUF1cEhFcXFDcFV3aERQQmVvR29hRGNVYUZqaWpwRWloSzJxYlNtcmpLUm9jM1F5dmFJZDF0TThUNWoiLCJkYXRha2V5IjoiQVFFQkFIaHdtMFlhSVNKZVJ0Sm01bjFHNnVxZWVrWHVvWFhQZTVVRmNlOVJxOC8xNHdBQUFINHdmQVlKS29aSWh2Y05BUWNHb0c4d2JRSUJBREJvQmdrcWhraUc5dzBCQndFd0hnWUpZSVpJQVdVREJBRXVNQkVFREpnQ2RJY3pRS2xUZEo1c3NRSUJFSUE3OGNEMUdxNnB6dW5NanBjRC82eURGZ0ZOTUk5SFo4ZnlKUEtoNXZ6bWROb3hZalFxNkx6TDFXU3hpeFROQzl5ZVRSMjhubWV0WDRsdjZUQT0iLCJ2ZXJzaW9uIjoiMiIsInR5cGUiOiJEQVRBX0tFWSIsImV4cGlyYXRpb24iOjE3MDE5NDYwMjR9