# Table: google_analytics_accounts

The primary key for this table is **id**.

## Relations

The following tables depend on google_analytics_accounts:
  - [google_analytics_web_properties](google_analytics_web_properties.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|child_link|JSON|
|created|String|
|id (PK)|String|
|kind|String|
|name|String|
|permissions|JSON|
|self_link|String|
|starred|Bool|
|updated|String|