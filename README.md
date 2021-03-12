# gun-lite
Same gRPC Tunnel, but without gRPC / Protobuf

## What's This?
Almost same thing as https://github.com/Qv2ray/gun, but without gRPC or Protobuf.

## Why no gRPC / no Protobuf?
Said @Dreamacro from https://github.com/Dreamacro/clash/pull/1287#issuecomment-797200851:

> gRPC (protobuf) too heavy-weight.
> Take linux-amd64 as an example. On dev branch, the binary size is 8.8M, but on this PR, the binary size is 13M. Increased size by nearly 50% new_moon_with_face

## How are you going to do it without gRPC / Protobuf?
We are going to reuse existing HTTP/2 libraries and simulate our own gRPC / Protobuf behavior.

We hope that this will greatly reduce the resulting binary size.

## Can I use this freely in my product?
Yes. Just follow the MIT License.
