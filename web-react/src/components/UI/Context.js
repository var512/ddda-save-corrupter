import { createContext } from 'react';

const UiContext = createContext({
  hasUserfile: {},
  setHasUserfile: () => {},
});

export default UiContext;
