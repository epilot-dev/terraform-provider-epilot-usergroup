terraform {
  required_providers {
    epilot-usergroup = {
      source  = "epilot-dev/epilot-usergroup"
      version = "0.10.3"
    }
  }
}

provider "epilot-usergroup" {
  # Configuration options

  epilot_auth="eyJraWQiOiJyd1wvSUdNeXJIU1wvV1wvMHlkOWoyT3Z6eURxTDEwM1RCUUNaUkVBejVMSUpBPSIsImFsZyI6IlJTMjU2In0.eyJzdWIiOiI3YmNjMzE2Ni03YmRhLTRmN2EtOTdjZi1mYjgxODAxYjUzNWUiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiaXNzIjoiaHR0cHM6XC9cL2NvZ25pdG8taWRwLmV1LWNlbnRyYWwtMS5hbWF6b25hd3MuY29tXC9ldS1jZW50cmFsLTFfNGh3VXlSWkVCIiwiY3VzdG9tOml2eV9vcmdfaWQiOiI3MzkyMjQiLCJjb2duaXRvOnVzZXJuYW1lIjoibi5nb2VsQGVwaWxvdC5jbG91ZCIsImN1c3RvbTppdnlfdXNlcl9pZCI6IjExMDAwMDM3IiwiYXVkIjoiNDdwcjdzdDdsNHVtYm1wZmJpanY2N21odWEiLCJldmVudF9pZCI6Ijk0ODg4YWIxLTJjOTktNDk1NC1hMzdkLWRlNjRhYTE1Nzk5NiIsInRva2VuX3VzZSI6ImlkIiwiYXV0aF90aW1lIjoxNzQwNzM0ODAyLCJleHAiOjE3NDEwOTUzODksImlhdCI6MTc0MTA5MTc4OSwiZW1haWwiOiJuLmdvZWxAZXBpbG90LmNsb3VkIn0.dEuWIz6EYuBHHAhgWgQ04ndX_4Mg33jMUdnCrIN9YXx0X1QvdRgvKD6VoXzqxXl23M1SlEt2z639l4WTx4h1Wfpf65EhXc6zPb-7UHTV40noV2z4tvODLb9Wn4TTpDWWqcT1xX5i18YZWOO527xRvkUJsyJulwseNoq18uKPqcrNNfLV1_3doDFbLLJvd__Ie1yJ__f4ZQyaGHu89WSGZG2_DJhzGeCiyZl2JK7Lzk7qYo2zx636zSMdRL-ti0NntwRxBkmvfzqUfXWhVf9u8qk6IBhL7XBTFnsBd1Tj9GfXrc3V3JHj4PFdBRRPmG2shL6PAlyVPQeInY-SVKZmJw"
  server_url = "https://user.dev.sls.epilot.io"
}

# import {
#   to = epilot-usergroup_user_group.my_usergroup
#   id= "v7s3qq2g2nxr5n4"
# }

resource "epilot-usergroup_user_group" "my_usergroup" {
  name = "Terraform User Group"
  user_ids = [
    "1234"
  ]
}


# resource "epilot-usergroup_user_group" "my_usergroup" {

# }