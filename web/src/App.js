import './App.css';
import './BootstrapTheme.css';
import UiContext from 'components/UI/Context';
import { MemoryRouter as Router } from 'react-router-dom';
import { useState } from 'react';
import Sidebar from 'components/Sidebar/Sidebar';
import Content from './components/Content/Content';

const App = () => {
  const [hasUserfile, setHasUserfile] = useState(false);
  const providerValue = { hasUserfile, setHasUserfile };

  return (
    <UiContext.Provider value={providerValue}>
      <Router>
        <div className="container-fluid">
          <div className="row">
            <Sidebar />
            <Content />
          </div>
        </div>
      </Router>
    </UiContext.Provider>
  );
};

export default App;
