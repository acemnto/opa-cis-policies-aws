package cis

default allow = false

allow {
  input.check_id == "1.1"
  input.evidence.root_account_mfa_enabled == true
}

allow {
  input.check_id == "1.2"
  pw := input.evidence.password_policy
  pw.minimum_length >= 14
  pw.require_symbols
  pw.require_numbers
  pw.require_uppercase_characters
  pw.require_lowercase_characters
}

allow {
  input.check_id == "2.1"
  input.evidence.cloudtrail_enabled_in_all_regions == true
}

allow {
  input.check_id == "3.1"
  pb := input.evidence.public_access_block
  pb.block_public_acls
  pb.block_public_policy
}

allow {
  input.check_id == "4.1"
  not input.evidence.contains_ssh_open_to_world
}
