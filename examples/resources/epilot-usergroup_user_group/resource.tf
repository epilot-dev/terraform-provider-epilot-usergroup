resource "epilot-usergroup_user_group" "my_usergroup" {
  image_uri = {
    additional_properties = { "gradient_colors" : ["#0588f0", "#3358d4"] }
    gradient_colors = [
      "#0588f0",
      "#3358d4",
    ]
    original     = "https://account-profile-images.epilot.cloud/org/group-avatar.png"
    thumbnail_32 = "https://account-profile-images.epilot.cloud/org/group-avatar_32x32.png"
    thumbnail_64 = "https://account-profile-images.epilot.cloud/org/group-avatar_64x64.png"
  }
  name = "Finance"
}