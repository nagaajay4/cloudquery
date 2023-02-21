# Table: google_analytics_profiles

The composite primary key for this table is (**account_id**, **id**, **web_property_id**).

## Relations

This table depends on [google_analytics_web_properties](google_analytics_web_properties.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|bot_filtering_enabled|Bool|
|child_link|JSON|
|created|String|
|currency|String|
|default_page|String|
|e_commerce_tracking|Bool|
|enhanced_e_commerce_tracking|Bool|
|exclude_query_parameters|String|
|id (PK)|String|
|internal_web_property_id|String|
|kind|String|
|name|String|
|parent_link|JSON|
|permissions|JSON|
|self_link|String|
|site_search_category_parameters|String|
|site_search_query_parameters|String|
|starred|Bool|
|strip_site_search_category_parameters|Bool|
|strip_site_search_query_parameters|Bool|
|timezone|String|
|type|String|
|updated|String|
|web_property_id (PK)|String|
|website_url|String|