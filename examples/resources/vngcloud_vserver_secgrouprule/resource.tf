resource "vngcloud_vserver_secgrouprule" "secgrouprule" {
  project_id        = "pro-462803f3-6858-466f-bf05-df2b33faa360"
  direction         = "ingress"
  ethertype         = "IPv4"
  port_range_max    = 65535
  port_range_min    = count.index
  protocol          = "TCP"
  remote_ip_prefix  = "0.0.0.0/0"
  security_group_id = "secg-ce0669a5-28c8-4263-8a0a-01adc93fc5a3"
}