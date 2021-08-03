import { Table } from 'react-bootstrap';
import { BACKEND_URL } from 'constants/app';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

const AppearanceOverride = (props) => (
  <div id="pawn-appearance-override">
    <h2>Appearance override</h2>
    <p>Changes made to the pawn (with precedence).</p>
    <Table striped bordered>
      <tbody>
      <tr>
        <th>
          Actions
        </th>
        <td>
          <a href={`${BACKEND_URL}/pawns/main/export-with-appearance-override`} className="btn btn-primary btn-sm"
             target="_blank" rel="noreferrer">
            <FontAwesomeIcon icon={'download'} /> Export .xml with appearance override
          </a>
        </td>
      </tr>
      <tr>
        <th>
          Name
        </th>
        <td>
          {props.data.pawn.data.AppearanceOverride.Attributes.Name}
        </td>
      </tr>
      <tr>
        <th>
          Gender
        </th>
        <td>
          {props.data.pawn.data.AppearanceOverride.Attributes.Gender.Value === 0 ? 'Male' : 'Female'}
        </td>
      </tr>
      <tr>
        <th>
          Nickname
        </th>
        <td>
          {props.data.pawn.data.AppearanceOverride.Attributes.Nickname.Value}
        </td>
      </tr>
      </tbody>
    </Table>

    <pre className="mt-3">
      {props.data.pawn.data.AppearanceOverride.DataXML}
    </pre>

    <hr className="mt-5 mb-5" />

  </div>
);

export default AppearanceOverride;
