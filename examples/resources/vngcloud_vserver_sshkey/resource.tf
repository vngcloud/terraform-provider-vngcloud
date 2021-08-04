resource "vngcloud_vserver_sshkey" "sshkey" {
    project_id = "pro-462803f3-6858-466f-bf05-df2b33faa360"
    name = "sshkey_name"
    public_key = "ssh_public_key"
    lifecycle {
        create_before_destroy = false
    }
}