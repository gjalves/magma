/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import type {CheckListItem} from '../checkListCategory/ChecklistItemsDialogMutateState';

import FormField from '@fbcnms/ui/components/design-system/FormField/FormField';
import React from 'react';
import Text from '@fbcnms/ui/components/design-system/Text';
import TextInput from '@fbcnms/ui/components/design-system/Input/TextInput';
import {makeStyles} from '@material-ui/styles';

type Props = {
  item: CheckListItem,
  onChange?: (updatedChecklistItem: CheckListItem) => void,
};

const useStyles = makeStyles(() => ({
  container: {
    display: 'flex',
    flexDirection: 'row',
    alignItems: 'center',
  },
  expandindPart: {
    flexGrow: 1,
    flexBasis: 0,
    '&:not(:first-child)': {
      marginLeft: '8px',
    },
    '&:not(:last-child)': {
      marginRight: '8px',
    },
  },
}));

const FreeTextCheckListItemFilling = ({item, onChange}: Props) => {
  const classes = useStyles();

  const _updateOnChange = newValue => {
    if (!onChange) {
      return;
    }
    const updatedItem = {
      ...item,
      stringValue: newValue,
      checked: !!newValue && newValue.trim().length > 0,
    };
    onChange(updatedItem);
  };

  return (
    <div className={classes.container}>
      <Text className={classes.expandindPart} variant="body2" weight="regular">
        {item.title}
      </Text>
      <FormField>
        <TextInput
          className={classes.expandindPart}
          type="string"
          placeholder={item.helpText || ''}
          value={item.stringValue || ''}
          onChange={event => _updateOnChange(event.target.value)}
        />
      </FormField>
    </div>
  );
};

export default FreeTextCheckListItemFilling;
