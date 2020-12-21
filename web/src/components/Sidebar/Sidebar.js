import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { NavLink } from 'react-router-dom';
import { useContext } from 'react';
import UiContext from 'components/UI/Context';

const Sidebar = () => {
  const { hasUserfile } = useContext(UiContext);

  return (
    <nav id="sidebar-menu" className="col-md-4 col-lg-3 col-xl-3 d-md-block sidebar">
      <div className="sidebar-brand">
        <span>DDDA Save Corrupter</span>
      </div>
      <ul className="nav nav-tabs flex-column">
        <li className="nav-item">
          <NavLink to="/import-file" className="nav-link" activeClassName="active">
            <FontAwesomeIcon icon={'file'} />
            Import file
          </NavLink>
        </li>
        <li className={hasUserfile ? 'nav-item' : 'd-none'}>
          <NavLink to="/pawns/main" className="nav-link" activeClassName="active">
            <FontAwesomeIcon icon={'chess-knight'} />
            Main Pawn
          </NavLink>
        </li>
        <li className={hasUserfile ? 'nav-item' : 'd-none'}>
          <NavLink to="/pawns/first" className="nav-link" activeClassName="active">
            <FontAwesomeIcon icon={'chess-pawn'} />
            First Pawn
          </NavLink>
        </li>
        <li className={hasUserfile ? 'nav-item' : 'd-none'}>
          <NavLink to="/pawns/second" className="nav-link" activeClassName="active">
            <FontAwesomeIcon icon={'chess-pawn'} />
            Second Pawn
          </NavLink>
        </li>
        <li className={hasUserfile ? 'nav-item' : 'd-none'}>
          <NavLink to="/export-file" className="nav-link" activeClassName="active">
            <FontAwesomeIcon icon={'download'} />
            Export file
          </NavLink>
        </li>
      </ul>
    </nav>
  );
};

export default Sidebar;
