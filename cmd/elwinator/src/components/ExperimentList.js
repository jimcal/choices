// @flow
import React from 'react';
import { Link } from 'react-router';
import { connect } from 'react-redux';

import { experimentURL } from '../urls';
import { experimentDelete } from '../actions';
import { getNamespace } from '../reducers/namespaces';
import { getExperiments } from '../reducers/experiments';
import { getParams } from '../reducers/params';

const ExperimentList = ({ namespaceName, experiments, dispatch }) => {
  const exps = experiments.map((e, i) =>
    <tr key={e.name}>
      <td>{i + 1}</td>
      <td><Link to={experimentURL(e.id)}>{e.name}</Link></td>
      <td>{e.params.map(p => p.name).join(', ')}</td>
      <td>
        <button
          className="btn btn-default btn-xs"
          onClick={() => dispatch(experimentDelete(namespaceName, e.id))}
        >&times;</button>
      </td>
    </tr>
  );
  return (
    <table className="table table-striped">
    <thead>
      <tr>
        <th>#</th>
        <th>Experiment</th>
        <th>Param(s)</th>
        <th>Delete</th>
      </tr>
    </thead>
    <tbody>
      {exps}
    </tbody>
    </table>
  );
}

ExperimentList.PropTypes = {
  namespaceName: React.PropTypes.string.isRequired,
}

const mapStateToProps = (state, ownProps) => {
  const ns = getNamespace(state.entities.namespaces, ownProps.namespaceName);
  const experiments = getExperiments(state.entities.experiments, ns.experiments)
  .map(e => ({
    ...e,
    params: getParams(state.entities.params, e.params),
  }));
  return {
    experiments,
  }
}

const connected = connect(mapStateToProps)(ExperimentList);

export default connected;
