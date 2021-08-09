resource "vngcloud_vdb_configuration_group" "config_group" {
  datastore_type    = "MySQL"
  datastore_version = "5.7"
  name              = "test configuration group"
}