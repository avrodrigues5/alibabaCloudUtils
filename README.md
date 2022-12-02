# alibabaCloudUtils

To configure Billing to an Object Storage Service using your Access Key and Access Secret you could use this code

```go run main.go <ACCESS_KEY> <ACCESS_SECRET> <REGION_ID> <OSS_BUCKET_NAME> <SUBSCRIPTION_TYPE>```

ACCESS_KEY and ACCESS_SECRET is created in Alibaba Cloud Console > Access Key Management
REGION_ID is the region associated with the OSS_BUCKET
OSS_BUCKET_NAME is the name of the OSS Bucket
SUBSCRIPTION_TYPE Can be either `InstanceDetailForBillingPeriod` or `BillingItemDetailForBillingPeriod`