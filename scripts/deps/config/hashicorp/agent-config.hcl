exit_after_auth = false

vault {
  address = "http://vault:8200"
}

auto_auth {
  method "approle" {
    config = {
      role_id_file_path = "/vault/token/role"
      secret_id_file_path = "/vault/token/secret"
      remove_secret_id_file_after_reading = false
    }
  }

  sink "file" {
    wrap_ttl = "30s"
    config = {
      path = "/vault/token/.vault-token"
    }
  }
}
