// @flow
import { v4 } from 'node-uuid';

/**
 * entitiesLoaded is an action that merges all the entities from the elwin storage to the store.
 * @param {Object} entities - entities object.
 */
export const entitiesLoaded = (entities: any) => ({
  type: 'ENTITIES_LOADED',
  entities,
});

/**
 * changesClear is an action that clears the changes listed from the change list.
 * @param {Array} changes - list of namespaces to remove.
 */
export const changesClear = (changes: Array<string>) => ({
  type: 'CHANGES_CLEAR',
  changes,
});

/**
 * namespaceAdd is an action that adds a namespace to the namespace list.
 * @param {string} name - The name of the namespace.
 */
export const namespaceAdd = (namespace: string) => ({
  type: 'NAMESPACE_ADD',
  id: v4(),
  namespace,
});

/**
 * namespaceDelete is an action that marks a namespace to be deleted.
 * @param {string} namespace - The namespace to delete.
 */
export const namespaceDelete = (namespace: string) => ({
  type: 'NAMESPACE_DELETE',
  namespace,
});

/**
 * namespaceLocalDelete is an action that removes the given namespace from the
 * local state. This would be used when you create a new namespace don't
 * publish it then decide to delete it.
 * @param {string} namespace - The namespace you want to delete.
 */
export const namespaceLocalDelete = (namespace: string) => ({
  type: 'NAMESPACE_LOCAL_DELETE',
  namespace,
});

/**
 * namespaceName is an action that set the name of the namespace.
 * @param {string} namespace - The original name of the namespace.
 * @param {string} name - The new name of the namespace.
 */
export const namespaceName = (namespace: string, name: string) => ({
  type: 'NAMESPACE_NAME',
  namespace,
  name,
});

/**
 * namespaceToggleLabel is an action that toggles the specified label.
 * @param {string} namespace - name of the namespace the label is in.
 * @param {string} id - id of the label to toggle.
 */
export const namespaceToggleLabel = (namespace: string, id: string) => ({
  type: 'NAMESPACE_TOGGLE_LABEL',
  namespace,
  id,
})

/**
 * namespaceAddLabel is an action that adds a label to an experiment.
 * @param {string} namespace - The namespace for the label.
 * @param {string} name - The name of the label to add.
 */
export const namespaceAddLabel = (namespace: string, name: string) => ({
  type: 'NAMESPACE_ADD_LABEL',
  id: v4(),
  namespace,
  name,
});

/**
 * namespaceTogglePublish is an action that toggles a namespace for publishing.
 * @param {string} namespace - The namespace to publish.
 */
export const namespaceTogglePublish = (namespace: string) => ({
  type: 'NAMESPACE_TOGGLE_PUBLISH',
  namespace,
});

/** 
 * experimentAdd is an action that adds and experiment to the namespace.
 * @param {string} namespace - The namespace for the experiment.
 * @param {string} name - The name of the experiment to add to the namespace.
 */
export const experimentAdd = (namespace: string, name: string) => ({
  type: 'EXPERIMENT_ADD',
  id: v4(),
  namespace,
  name,
});

/**
 * experimentDelete is an action that marks an experiment for deletion.
 * @param {string} namespace - The namespace the experiment is in.
 * @param {string} experiment - The name of the experiment to delete.
 */
export const experimentDelete = (namespace: string, experiment: string) => ({
  type: 'EXPERIMENT_DELETE',
  namespace,
  experiment,
});

/**
 * experimentName is an action that sets the name in an experiment.
 * @param {string} namespace - The experiment namespace.
 * @param {string} experiment - The experiment's original name.
 * @param {string} name - The experiment's new name.
 */
export const experimentName = (namespace: string, experiment: string, name: string) => ({
  type: 'EXPERIMENT_NAME',
  namespace,
  experiment,
  name,
});

/**
 * experimentNumSegments is an action that sets the number of segments in an
 * expermient.
 * @param {string} namespace - The experiment namespace.
 * @param {string} experiment - The experiment that is being changed.
 * @param {Array} namespaceSegments - The segments claimed by the namespace.
 * @param {number} numSegments - The number of segments the experiment
 * should have.
 */
export const experimentNumSegments = (namespace: string, experiment: string, namespaceSegments: Uint8Array, numSegments: number) => ({
  type: 'EXPERIMENT_NUM_SEGMENTS',
  namespace,
  experiment,
  namespaceSegments,
  numSegments,
});

/**
 * paramName is an action that sets the param name in an experiments param.
 * @param {string} namespace - The param namespace.
 * @param {string} param - The param's id.
 * @param {string} name - The param's new name.
 */
export const paramName = (namespace: string, param: string, name: string) => ({
  type: 'PARAM_NAME',
  namespace,
  param,
  name,
});

/**
 * paramAdd is an action that adds a param to an experiment.
 * @param {string} namespace - The namespace name.
 * @param {string} experiment - The experiment id.
 * @param {Object} param - The param you are adding.
 */
export const paramAdd = (namespace: string, experiment: string, name:string) => ({
  type: 'PARAM_ADD',
  id: v4(),
  namespace,
  experiment,
  name,
});

/**
 * paramDelete is an action that deletes a param from an experiment.
 * @param {string} namespace - The namespace that the param is in.
 * @param {string} experiment - The experiment that the param is in.
 * @param {string} name - The name of the param to delete.
 */
export const paramDelete = (namespace: string, experiment: string, param: string) => ({
  type: 'PARAM_DELETE',
  namespace,
  experiment,
  param,
});

/**
 * paramToggleWeighted is an action that toggles whether a param is weighted or
 * uniform.
 * @param {string} namespace - The param namespace.
 * @param {string} param - The name of the param.
 */
export const paramToggleWeighted = (namespace: string, param: string) => ({
  type: 'PARAM_TOGGLE_WEIGHTED',
  param,
});

/**
 * paramAddChoice is an action that adds a choice to a param. You must also call
 * paramAddWeight if the param is a weighted param.
 * @param {string} namespace - The namespace that the param is in.
 * @param {string} experiment - The experiment's name.
 * @param {string} param - The name of the param.
 * @param {string} choice - The choice to add to the param.
 */
export const paramAddChoice = (namespace: string, param: string, choice: string) => ({
  type: 'PARAM_ADD_CHOICE',
  namespace,
  param,
  choice,
});

/**
 * paramDeleteChoice is an action that deletes a choice from a param.
 * @param {string} namespace - The param namespace.
 * @param {string} param - The param name.
 * @param {string} index - The index of the choice to delete.
 */
export const paramDeleteChoice = (namespace: string, param: string, index: number) => ({
  type: 'PARAM_DELETE_CHOICE',
  namespace,
  param,
  index,
});

/**
 * paramAddWeight is an action that adds a weight to a param. You should always
 * call paramAddChoice before calling this.
 * @param {string} namespace - The param namespace.
 * @param {string} param - The name of the param.
 * @param {string} weight - The weight to add to the param.
 */
export const paramAddWeight = (namespace: string, param: string, weight: number) => ({
  type: 'PARAM_ADD_WEIGHT',
  namespace,
  param,
  weight,
});

/**
 * paramClearChoices is an action that removes all the choices and weights.
 * @param {string} namespace - The namespace.
 * @param {string} param - The name of the param.
 */
export const paramClearChoices = (namespace: string, param: string) => ({
  type: 'PARAM_CLEAR_CHOICES',
  namespace,
  param,
});
