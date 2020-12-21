import { Table } from 'react-bootstrap';
import { BACKEND_URL } from 'constants/app';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { Link } from 'react-router-dom';
import AppearanceOverride from 'components/Content/Pawns/AppearanceOverride';

const Information = (props) => (
  <div>
    {
      props.data.category === 'main'
        ? <AppearanceOverride data={{ pawn: props.data.pawn }} />
        : null
    }
    <div id="pawn-information">
      <h2>Pawn information</h2>
      <Table striped bordered>
        <tbody>
        <tr>
          <th>
            Actions
          </th>
          <td>
            <Link to={`/pawns/${props.data.category}/replace`} className="btn btn-primary btn-sm mr-2">
              <FontAwesomeIcon icon={'people-arrows'} /> Replace
            </Link>
            <a href={`${BACKEND_URL}/pawns/${props.data.category}/export`} className="btn btn-primary btn-sm"
               target="_blank" rel="noreferrer">
              <FontAwesomeIcon icon={'download'} /> Export .xml
            </a>
          </td>
        </tr>
        <tr>
          <th>
            Name
          </th>
          <td>
            {props.data.pawn.data.Attributes.Name}
          </td>
        </tr>
        <tr>
          <th>
            Gender
          </th>
          <td>
            {props.data.pawn.data.Attributes.Gender.Value === 0 ? 'Male' : 'Female'}
          </td>
        </tr>
        <tr>
          <th>
            Nickname
          </th>
          <td>
            {props.data.pawn.data.Attributes.Nickname.Value}
          </td>
        </tr>
        </tbody>
      </Table>

      <pre className="mt-3">
        <code className="language-css">
          {props.data.pawn.dataXML}
        </code>
      </pre>
    </div>
  </div>
);

export default Information;
