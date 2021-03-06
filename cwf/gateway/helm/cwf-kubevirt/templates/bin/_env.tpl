# Copyright 2020 The Magma Authors.

# This source code is licensed under the BSD-style license found in the
# LICENSE file in the root directory of this source tree.

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

COMPOSE_PROJECT_NAME=cwf
DOCKER_REGISTRY={{ .Values.cwf.image.docker_registry }}
DOCKER_USERNAME={{ .Values.cwf.image.username }}
DOCKER_PASSWORD={{ .Values.cwf.image.password }}
IMAGE_VERSION={{ .Values.cwf.image.tag }}
BUILD_CONTEXT={{ .Values.cwf.repo.url }}#{{ .Values.cwf.repo.branch }}

ROOTCA_PATH=/var/opt/magma/certs/rootCA.pem
CONTROL_PROXY_PATH=/etc/magma/control_proxy.yml
CONFIGS_TEMPLATES_PATH=/etc/magma/templates

CERTS_VOLUME=/var/opt/magma/certs
CONFIGS_OVERRIDE_VOLUME=/var/opt/magma/configs
CONFIGS_DEFAULT_VOLUME=/etc/magma
SECRETS_VOLUME=/var/opt/magma/secrets

{{ if .Values.cwf.dpi }}
DPI_LICENSE_NAME={{ .Values.cwf.dpi.dpi_license_name }}
{{- end }}

{{ if .Values.cwf.env }}
{{ .Values.cwf.env }}
{{- end }}
