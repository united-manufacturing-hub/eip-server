# Copyright 2025 UMH Systems GmbH
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM golang:1.24

LABEL org.opencontainers.image.source=https://github.com/united-manufacturing-hub/eip-server

COPY . /go/src/github.com/united-manufacturing-hub/eip-server

WORKDIR /go/src/github.com/united-manufacturing-hub/eip-server

RUN CGO_ENABLED=0 go build -v -o /eip-server cmd/server/main.go

FROM scratch

COPY --from=0 /eip-server /eip-server

WORKDIR /

ENTRYPOINT ["/eip-server"]
