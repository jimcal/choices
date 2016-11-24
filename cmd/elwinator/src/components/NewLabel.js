import React, { PropTypes } from 'react';
import { browserHistory } from 'react-router';

import { namespaceURL } from '../urls';
import { addLabel } from '../actions';

const NewLabel = ({ namespaceName, redirectOnSubmit, dispatch }) => {
  let label;
  return (
    <form onSubmit={e => {
      e.preventDefault();
      if (!label.value.trim()) {
        return;
      }
      dispatch(addLabel(namespaceName, label.value));
      if (!redirectOnSubmit) {
        label.value = '';
        return;
      }
      browserHistory.push(namespaceURL(namespaceName));
    }}>
      <div className="form-group">
        <label>Create new label</label>
        <input type="text"
          className="form-control"
          placeholder="Enter a label name"
          ref={node => label = node}
        />
      </div>
      <button type="submit" className="btn btn-primary">Create label</button>
    </form>
  );
}

NewLabel.propTypes = {
  namespaceName: PropTypes.string.isRequired,
  redirectOnSubmit: PropTypes.bool.isRequired,
}

export default NewLabel;
