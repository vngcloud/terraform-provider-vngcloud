
data "vngcloud_vlb_lb_packages" "lb_packages" {
  project_id = var.project_id
}

resource "vngcloud_vlb_load_balancer" "example"{
  project_id = var.project_id
  name       = "example_vlb"
  package_id =  data.vngcloud_vlb_lb_packages.lb_packages.packages[0].uuid
  scheme     = "Internet"
  subnet_id  = var.subnet_id
  type       = "Layer 7"
}


resource "vngcloud_vlb_pool" "example" {
  project_id       = var.project_id
  load_balancer_id = vngcloud_vlb_load_balancer.example.id
  name             = "example_pool"
  protocol         = "HTTP"
  stickiness       = true
  algorithm        = "ROUND_ROBIN"
  tls_encryption   = true
  health_monitor {
    health_check_method   = "GET"
    health_check_path     = "/"
    health_check_protocol = "HTTP"
    healthy_threshold     = 3
    unhealthy_threshold   = 3
    interval              = 30
    timeout               = 5
    success_code          = 200
    http_version          = "1.0"
  }
  members {
    backup       = false
    ip_address   = "10.0.1.8"
    monitor_port = 80
    port         = 80
    weight       = 1
    name         = "member1"
  }
  members {
    backup       = true
    ip_address   = "10.0.1.7"
    monitor_port = 81
    port         = 81
    weight       = 1
    name         = "member2"
  }
  members {
    backup       = false
    ip_address   = "10.0.1.6"
    monitor_port = 82
    port         = 82
    weight       = 1
    name         = "member3"
  }
}

resource "vngcloud_vlb_listener" "example" {
  project_id                    = var.project_id
  load_balancer_id              = vngcloud_vlb_load_balancer.example.id
  name                          = "example_listener"
  allowed_cidrs                 = "0.0.0.0/0"
  protocol                      = "HTTP"
  protocol_port                 = 81
  timeout_client                = 50
  timeout_connection            = 5
  timeout_member                = 60
  default_pool_id               = vngcloud_vlb_pool.example.id
}

resource "vngcloud_vlb_l7policy" "example" {
  project_id         = var.project_id
  load_balancer_id   = vngcloud_vlb_load_balancer.example.id
  listener_id        = vngcloud_vlb_listener.example.id
  name               = "policyLB7"
  rule_type          = "PATH"
  compare_type       = "EQUAL_TO"
  rule_value         = "https://www.example.com"
  action             = "REDIRECT_TO_URL"
  redirect_http_code = 302
  redirect_url       = "https://www.example.com"
  keep_query_string  = true
}
