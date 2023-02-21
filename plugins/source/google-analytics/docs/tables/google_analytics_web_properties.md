# Table: google_analytics_web_properties

The composite primary key for this table is (**account_id**, **id**).

## Relations

This table depends on [google_analytics_accounts](google_analytics_accounts.md).

The following tables depend on google_analytics_web_properties:
  - [google_analytics_profiles](google_analytics_profiles.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|child_link|JSON|
|created|String|
|data_retention_reset_on_new_activity|Bool|
|data_retention_ttl|String|
|default_profile_id|Int|
|id (PK)|String|
|industry_vertical|String|
|internal_web_property_id|String|
|kind|String|
|level|String|
|name|String|
|parent_link|JSON|
|permissions|JSON|
|profile_count|Int|
|self_link|String|
|starred|Bool|
|updated|String|
|website_url|String|