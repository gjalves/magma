/**
 * Copyright 2020 The Magma Authors.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @flow strict-local
 * @format
 */

import type {WithStyles} from '@material-ui/core';

import React from 'react';
import Text from './design-system/Text';
import {gray13} from '../theme/colors';
import {withStyles} from '@material-ui/core/styles';

type Props = {
  title: string,
  subtitle: ?string,
  className?: string,
} & WithStyles<typeof styles>;

const styles = _theme => ({
  title: {
    display: 'block',
  },
  subtitle: {
    color: gray13,
  },
});

const ConfigureTitle = (props: Props) => {
  const {title, subtitle, classes, className} = props;
  return (
    <div className={className}>
      <Text className={classes.title} variant="h6">
        {title}
      </Text>
      <Text className={classes.subtitle} variant="subtitle2">
        {subtitle}
      </Text>
    </div>
  );
};

export default withStyles(styles)(ConfigureTitle);
