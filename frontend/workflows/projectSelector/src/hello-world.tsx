import * as React from "react";
import { TextField } from "@clutch-sh/core";
import styled from "@emotion/styled";
import { Divider, LinearProgress } from "@material-ui/core";
import LayersIcon from "@material-ui/icons/Layers";
import _ from "lodash";

import ProjectGroup from "./project-group";
import { selectorReducer } from "./selector-reducer";

export enum Group {
  PROJECTS,
  UPSTREAM,
  DOWNSTREAM,
}

type UserActionKind =
  | "ADD_PROJECTS"
  | "REMOVE_PROJECTS"
  | "TOGGLE_PROJECTS"
  | "TOGGLE_ENTIRE_GROUP"
  | "ONLY_PROJECTS";

interface UserAction {
  type: UserActionKind;
  payload: UserPayload;
}

interface UserPayload {
  group: Group;
  projects?: string[];
}

type BackgroundActionKind = "HYDRATE_START" | "HYDRATE_END";

interface BackgroundAction {
  type: BackgroundActionKind;
  payload?: BackgroundPayload;
}

interface BackgroundPayload {
  result: any;
}

export type Action = BackgroundAction | UserAction;

export interface State {
  [Group.PROJECTS]: GroupState;
  [Group.UPSTREAM]: GroupState;
  [Group.DOWNSTREAM]: GroupState;

  projectData: { [projectName: string]: Project };
  loading: boolean;
}

// TODO: subout with full manifest structure (from proto def)
interface Project {
  upstreams: string[];
  downstreams: string[];
}

interface GroupState {
  [projectName: string]: ProjectState;
}

interface ProjectState {
  checked: boolean;
  // TODO: hidden should be derived?
  hidden?: boolean; // upstreams and downstreams are hidden when their parent is unchecked unless other parents also use them.
  custom?: boolean;
}

const StateContext = React.createContext<State | undefined>(undefined);
export const useReducerState = () => {
  return React.useContext(StateContext);
};

const DispatchContext = React.createContext<(action: Action) => void | undefined>(undefined);
export const useDispatch = () => {
  return React.useContext(DispatchContext);
};

const fakeAPI = (state: State) => {
  return {
    clutch: {
      upstreams: ["rides", "locations"],
      downstreams: ["queueworker"],
    },
  };
};

// TODO(perf): call with useMemo().
export const deriveSwitchStatus = (state: State, group: Group): boolean => {
  return (
    Object.keys(state[group]).length > 0 &&
    Object.keys(state[group]).every(key => state[group][key].checked)
  );
};

const initialState: State = {
  [Group.PROJECTS]: {},
  [Group.UPSTREAM]: {},
  [Group.DOWNSTREAM]: {},
  projectData: {},
  loading: false,
};

const SelectorContainer = styled.div({
  backgroundColor: "#F9FAFE",
  border: "1px solid rgba(13, 16, 48, 0.1)",
  boxShadow: "0px 4px 6px rgba(53, 72, 212, 0.2)",
  width: "245px",
  padding: "16px",
});

// TODO: change icon, center align icon and title
const StyledWorkflowHeader = styled.div({
  paddingBottom: "16px",
});

const StyledWorkflowTitle = styled.span({
  fontWeight: "bold",
  fontSize: "20px",
  lineHeight: "24px",
  margin: "0px 8px",
});

// TODO: add plus icon in the text field
const StyledProjectTextField = styled.div({
  paddingTop: "16px",
});

const ProjectSelector = () => {
  // On load, we'll request a list of owned projects and their upstreams and downstreams from the API.
  // The API will contain information about the relationships between projects and upstreams and downstreams.
  // By default, the owned projects will be checked and others will be unchecked.
  // If a project is unchecked, the upstream and downstreams related to it disappear from the list.
  // If a project is rechecked, the checks were preserved.

  const [customProject, setCustomProject] = React.useState("");

  const [state, dispatch] = React.useReducer(selectorReducer, initialState);

  React.useEffect(() => {
    console.log("effect");
    // Determine if any hydration is required.
    // - Are any services missing from state.projectdata?
    // - Are projects empty (first load)?
    // - Is loading not already in progress?

    let allPresent = true;
    _.forEach(Object.keys(state[Group.PROJECTS]), p => {
      if (!(p in state.projectData)) {
        allPresent = false;
        return false; // Stop iteration.
      }
      return true; // Continue.
    });

    if (!state.loading && (Object.keys(state[Group.PROJECTS]).length == 0 || !allPresent)) {
      console.log("calling API!", state.loading);
      dispatch({ type: "HYDRATE_START" });
      // TODO: call API and use payload.
      setTimeout(
        () => dispatch({ type: "HYDRATE_END", payload: { result: fakeAPI(state) } }),
        1000
      );
    }
  }, [state[Group.PROJECTS]]);

  const handleAdd = () => {
    if (customProject === "") {
      return;
    }
    dispatch({
      type: "ADD_PROJECTS",
      payload: { group: Group.PROJECTS, projects: [customProject] },
    });
    setCustomProject("");
  };

  return (
    <DispatchContext.Provider value={dispatch}>
      <StateContext.Provider value={state}>
        <SelectorContainer>
          {state.loading && <LinearProgress color="secondary" />}
          <StyledWorkflowHeader>
            <LayersIcon />
            <StyledWorkflowTitle>Dash</StyledWorkflowTitle>
          </StyledWorkflowHeader>
          <Divider />
          <StyledProjectTextField>
            <TextField
              disabled={state.loading}
              placeholder="Add a project"
              value={customProject}
              onChange={e => setCustomProject(e.target.value)}
              onKeyDown={e => e.key === "Enter" && handleAdd()}
            />
          </StyledProjectTextField>
          <ProjectGroup title="Projects" group={Group.PROJECTS} />
          <Divider />
          <ProjectGroup title="Upstreams" group={Group.UPSTREAM} collapsible />
          <Divider />
          <ProjectGroup title="Downstreams" group={Group.DOWNSTREAM} collapsible />
        </SelectorContainer>
      </StateContext.Provider>
    </DispatchContext.Provider>
  );
};

const HelloWorld = () => (
  <>
    <ProjectSelector />
  </>
);

export default HelloWorld;
