/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow strict-local
 * @format
 */

import Card from '@material-ui/core/Card';
import EnodebKPIs from './EnodebKPIs';
import GatewayKPIs from './GatewayKPIs';
import Grid from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';
import React from 'react';
import Text from '../theme/design-system/Text';

import {CardTitleRow} from './layout/CardTitleRow';
import {GpsFixed} from '@material-ui/icons';
import {makeStyles} from '@material-ui/styles';

const useStyles = makeStyles(theme => ({
  eventsTable: {
    marginTop: theme.spacing(4),
    textAlign: 'center',
    padding: theme.spacing(10),
  },
}));

export default function () {
  const classes = useStyles();

  return (
    <>
      <CardTitleRow icon={GpsFixed} label="Events (388)" />
      <Grid container item zeroMinWidth alignItems="center" spacing={4}>
        <Grid item xs={12} md={6}>
          <Paper elevation={0}>
            <GatewayKPIs />
          </Paper>
        </Grid>
        <Grid item xs={12} md={6}>
          <Paper elevation={0}>
            <EnodebKPIs />
          </Paper>
        </Grid>
      </Grid>
      <Card elevation={0} className={classes.eventsTable}>
        <Text variant="body2">Events Table Goes Here</Text>
      </Card>
    </>
  );
}
