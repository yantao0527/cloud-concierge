package terraformerCLI

import terraformValueObjects "github.com/dragondrop-cloud/driftmitigation/implementations/terraform_value_objects"

var awsResourceGroups = map[terraformValueObjects.ResourceName]string{
	"aws_accessanalyzer_analyzer":             "accessanalyzer",
	"aws_acm_certificate":                     "acm",
	"aws_lb":                                  "alb",
	"aws_lb_listener":                         "alb",
	"aws_lb_listener_rule":                    "alb",
	"aws_lb_listener_certificate":             "alb",
	"aws_lb_target_group":                     "alb",
	"aws_lb_target_group_attachment":          "alb",
	"aws_api_gateway_authorizer":              "api_gateway",
	"aws_api_gateway_api_key":                 "api_gateway",
	"aws_api_gateway_documentation_part":      "api_gateway",
	"aws_api_gateway_gateway_response":        "api_gateway",
	"aws_api_gateway_integration":             "api_gateway",
	"aws_api_gateway_integration_response":    "api_gateway",
	"aws_api_gateway_method":                  "api_gateway",
	"aws_api_gateway_method_response":         "api_gateway",
	"aws_api_gateway_model":                   "api_gateway",
	"aws_api_gateway_resource":                "api_gateway",
	"aws_api_gateway_rest_api":                "api_gateway",
	"aws_api_gateway_stage":                   "api_gateway",
	"aws_api_gateway_usage_plan":              "api_gateway",
	"aws_api_gateway_vpc_link":                "api_gateway",
	"aws_appsync_graphql_api":                 "appsync",
	"aws_autoscaling_group":                   "auto_scaling",
	"aws_launch_configuration":                "auto_scaling",
	"aws_launch_template":                     "auto_scaling",
	"aws_batch_compute_environment":           "batch",
	"aws_batch_job_definition":                "batch",
	"aws_batch_job_queue":                     "batch",
	"aws_budgets_budget":                      "budgets",
	"aws_cloud9_environment_ec2":              "cloud9",
	"aws_cloudformation_stack":                "cloudformation",
	"aws_cloudformation_stack_set":            "cloudformation",
	"aws_cloudformation_stack_set_instance":   "cloudformation",
	"aws_cloudfront_distribution":             "cloudfront",
	"aws_cloudfront_cache_policy":             "cloudfront",
	"aws_cloudhsm_v2_cluster":                 "cloudhsm",
	"aws_cloudhsm_v2_hsm":                     "cloudhsm",
	"aws_cloudtrail":                          "cloudtrail",
	"aws_cloudwatch_dashboard":                "cloudwatch",
	"aws_cloudwatch_event_rule":               "cloudwatch",
	"aws_cloudwatch_event_target":             "cloudwatch",
	"aws_cloudwatch_metric_alarm":             "cloudwatch",
	"aws_codebuild_project":                   "codebuild",
	"aws_codecommit_repository":               "codecommit",
	"aws_codedeploy_app":                      "codedeploy",
	"aws_codepipeline":                        "codepipeline",
	"aws_codepipeline_webhook":                "codepipeline",
	"aws_cognito_identity_pool":               "cognito",
	"aws_cognito_user_pool":                   "cognito",
	"aws_config_config_rule":                  "config",
	"aws_config_configuration_recorder":       "config",
	"aws_config_delivery_channel":             "config",
	"aws_customer_gateway":                    "customer_gateway",
	"aws_datapipeline_pipeline":               "datapipeline",
	"aws_devicefarm_project":                  "devicefarm",
	"aws_docdb_cluster":                       "docdb",
	"aws_docdb_cluster_instance":              "docdb",
	"aws_docdb_cluster_parameter_group":       "docdb",
	"aws_docdb_subnet_group":                  "docdb",
	"aws_dynamodb_table":                      "dynamodb",
	"aws_ebs_volume":                          "ebs",
	"aws_volume_attachment":                   "ebs",
	"aws_instance":                            "ec2_instance",
	"aws_ecr_lifecycle_policy":                "ecr",
	"aws_ecr_repository":                      "ecr",
	"aws_ecr_repository_policy":               "ecr",
	"aws_ecrpublic_repository":                "ecrpublic",
	"aws_ecs_cluster":                         "ecs",
	"aws_ecs_service":                         "ecs",
	"aws_ecs_task_definition":                 "ecs",
	"aws_efs_access_point":                    "efs",
	"aws_efs_file_system":                     "efs",
	"aws_efs_file_system_policy":              "efs",
	"aws_efs_mount_target":                    "efs",
	"aws_eip":                                 "eip",
	"aws_eks_cluster":                         "eks",
	"aws_eks_node_group":                      "eks",
	"aws_elasticache_cluster":                 "elasticache",
	"aws_elasticache_parameter_group":         "elasticache",
	"aws_elasticache_subnet_group":            "elasticache",
	"aws_elasticache_replication_group":       "elasticache",
	"aws_elastic_beanstalk_application":       "elastic_beanstalk",
	"aws_elastic_beanstalk_environment":       "elastic_beanstalk",
	"aws_elb":                                 "elb",
	"aws_emr_cluster":                         "emr",
	"aws_emr_security_configuration":          "emr",
	"aws_network_interface":                   "eni",
	"aws_elasticsearch_domain":                "es",
	"aws_kinesis_firehose_delivery_stream":    "firehose",
	"aws_glue_crawler":                        "glue",
	"aws_glue_catalog_database":               "glue",
	"aws_glue_catalog_table":                  "glue",
	"aws_glue_job":                            "glue",
	"aws_glue_trigger":                        "glue",
	"aws_iam_access_key":                      "iam",
	"aws_iam_group":                           "iam",
	"aws_iam_group_policy":                    "iam",
	"aws_iam_group_policy_attachment":         "iam",
	"aws_iam_instance_profile":                "iam",
	"aws_iam_policy":                          "iam",
	"aws_iam_role":                            "iam",
	"aws_iam_role_policy":                     "iam",
	"aws_iam_role_policy_attachment":          "iam",
	"aws_iam_user":                            "iam",
	"aws_iam_user_group_membership":           "iam",
	"aws_iam_user_policy":                     "iam",
	"aws_iam_user_policy_attachment":          "iam",
	"aws_internet_gateway":                    "igw",
	"aws_iot_thing":                           "iot",
	"aws_iot_thing_type":                      "iot",
	"aws_iot_topic_rule":                      "iot",
	"aws_iot_role_alias":                      "iot",
	"aws_kinesis_stream":                      "kinesis",
	"aws_kms_key":                             "kms",
	"aws_kms_alias":                           "kms",
	"aws_kms_grant":                           "kms",
	"aws_lambda_event_source_mapping":         "lambda",
	"aws_lambda_function":                     "lambda",
	"aws_lambda_function_event_invoke_config": "lambda",
	"aws_lambda_layer_version":                "lambda",
	"aws_lambda_permission":                   "lambda",
	"aws_cloudwatch_log_group":                "logs",
	"aws_media_package_channel":               "media_package",
	"aws_media_store_container":               "media_store",
	"aws_medialive_channel":                   "medialive",
	"aws_medialive_input":                     "medialive",
	"aws_medialive_input_security_group":      "medialive",
	"aws_msk_cluster":                         "msk",
	"aws_network_acl":                         "nacl",
	"aws_nat_gateway":                         "nat",
	"aws_opsworks_application":                "opsworks",
	"aws_opsworks_custom_layer":               "opsworks",
	"aws_opsworks_instance":                   "opsworks",
	"aws_opsworks_java_app_layer":             "opsworks",
	"aws_opsworks_php_app_layer":              "opsworks",
	"aws_opsworks_rds_db_instance":            "opsworks",
	"aws_opsworks_stack":                      "opsworks",
	"aws_opsworks_static_web_layer":           "opsworks",
	"aws_opsworks_user_profile":               "opsworks",
	"aws_organizations_account":               "organization",
	"aws_organizations_organization":          "organization",
	"aws_organizations_organizational_unit":   "organization",
	"aws_organizations_policy":                "organization",
	"aws_organizations_policy_attachment":     "organization",
	"aws_qldb_ledger":                         "qldb",
	"aws_db_instance":                         "rds",
	"aws_db_proxy":                            "rds",
	"aws_db_cluster":                          "rds",
	"aws_db_cluster_snapshot":                 "rds",
	"aws_db_parameter_group":                  "rds",
	"aws_db_snapshot":                         "rds",
	"aws_db_subnet_group":                     "rds",
	"aws_db_option_group":                     "rds",
	"aws_db_event_subscription":               "rds",
	"aws_rds_global_cluster":                  "rds",
	"aws_resourcegroups_group":                "resourcegroups",
	"aws_route53_zone":                        "route53",
	"aws_route53_record":                      "route53",
	"aws_route_table":                         "route_table",
	"aws_main_route_table_association":        "route_table",
	"aws_route_table_association":             "route_table",
	"aws_s3_bucket":                           "s3",
	"aws_secretsmanager_secret":               "secretsmanager",
	"aws_securityhub_account":                 "securityhub",
	"aws_securityhub_member":                  "securityhub",
	"aws_securityhub_standards_subscription":  "securityhub",
	"aws_servicecatalog_portfolio":            "servicecatalog",
	"aws_ses_configuration_set":               "ses",
	"aws_ses_domain_identity":                 "ses",
	"aws_ses_email_identity":                  "ses",
	"aws_ses_receipt_rule":                    "ses",
	"aws_ses_receipt_rule_set":                "ses",
	"aws_ses_template":                        "ses",
	"aws_sfn_activity":                        "sfn",
	"aws_sfn_state_machine":                   "sfn",
	"aws_security_group":                      "sg",
	"aws_security_group_rule":                 "sg",
	"aws_sns_topic":                           "sns",
	"aws_sns_topic_subscription":              "sns",
	"aws_sqs_queue":                           "sqs",
	"aws_ssm_parameter":                       "ssm",
	"aws_subnet":                              "subnet",
	"aws_swf_domain":                          "swf",
	"aws_ec2_transit_gateway_route_table":     "transit_gateway",
	"aws_ec2_transit_gateway_vpc_attachment":  "transit_gateway",
	"aws_vpc":                                 "vpc",
	"aws_vpc_peering_connection":              "vpc_peering",
	"aws_vpn_connection":                      "vpn_connection",
	"aws_vpn_gateway":                         "vpn_gateway",
	"aws_waf_byte_match_set":                  "waf",
	"aws_waf_geo_match_set":                   "waf",
	"aws_waf_ipset":                           "waf",
	"aws_waf_rate_based_rule":                 "waf",
	"aws_waf_regex_match_set":                 "waf",
	"aws_waf_regex_pattern_set":               "waf",
	"aws_waf_rule":                            "waf",
	"aws_waf_rule_group":                      "waf",
	"aws_waf_size_constraint_set":             "waf",
	"aws_waf_sql_injection_match_set":         "waf",
	"aws_waf_web_acl":                         "waf",
	"aws_waf_xss_match_set":                   "waf",
	"aws_wafregional_byte_match_set":          "waf_regional",
	"aws_wafregional_geo_match_set":           "waf_regional",
	"aws_wafregional_ipset":                   "waf_regional",
	"aws_wafregional_rate_based_rule":         "waf_regional",
	"aws_wafregional_regex_match_set":         "waf_regional",
	"aws_wafregional_regex_pattern_set":       "waf_regional",
	"aws_wafregional_rule":                    "waf_regional",
	"aws_wafregional_rule_group":              "waf_regional",
	"aws_wafregional_size_constraint_set":     "waf_regional",
	"aws_wafregional_sql_injection_match_set": "waf_regional",
	"aws_wafregional_web_acl":                 "waf_regional",
	"aws_wafregional_xss_match_set":           "waf_regional",
	"aws_wafv2_ip_set":                        "wafv2_regional",
	"aws_wafv2_regex_pattern_set":             "wafv2_regional",
	"aws_wafv2_rule_group":                    "wafv2_regional",
	"aws_wafv2_web_acl":                       "wafv2_regional",
	"aws_wafv2_web_acl_association":           "wafv2_regional",
	"aws_wafv2_web_acl_logging_configuration": "wafv2_regional",
	"aws_workspaces_directory":                "workspaces",
	"aws_workspaces_ip_group":                 "workspaces",
	"aws_workspaces_workspace":                "workspaces",
	"aws_xray_sampling_rule":                  "xray",
}