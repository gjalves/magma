/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow strict-local
 * @format
 */

import React from 'react';

import {colors} from '../../theme/default';
import {makeStyles} from '@material-ui/styles';

type Props = {
  isGrey: boolean,
  isActive: boolean,
};

const useStyles = makeStyles(theme => ({
  status: {
    width: '8px',
    height: '8px',
    borderRadius: '50%',
    display: 'inline-block',
    marginRight: theme.spacing(1),
    backgroundColor: props =>
      props.isGrey
        ? colors.primary.nobel
        : props.isActive
        ? colors.state.positive
        : colors.state.error,
  },
}));

export default function DeviceStatusCircle(props: Props) {
  const classes = useStyles(props);
  return <span className={classes.status} />;
}
