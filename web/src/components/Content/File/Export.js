import { BACKEND_URL } from 'constants/app';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { useState } from 'react';
import ModalBox from 'components/UI/ModalBox';

const Export = () => {
  const [modalShow, setModalShow] = useState(false);
  const [modalMessage, setModalMessage] = useState('');

  // needs a server side api and a hook
  // golang broken pipe
  const block = () => {
    setModalMessage('Generating data. Download your file before continuing...');
    setModalShow(true);
  };

  return (
    <div>
      <ModalBox show={modalShow} backdrop="static" keyboard={false} onHide={() => setModalShow(false)} message={modalMessage} />
      <h2>Export file</h2>
      <form className="row">
        <div className="col-12">
          Generate and download a modified <span className="text-black-50">.sav</span> or <span
          className="text-black-50">.xml</span>
        </div>
        <div className="col-12 mt-2">
          <a
            href={`${BACKEND_URL}/files/sav/export`}
            className="btn btn-primary btn-lg mr-2"
            target="_blank"
            rel="noreferrer"
            onClick={block}
          >
            <FontAwesomeIcon icon={'download'} /> Export .sav
          </a>
          <a
            href={`${BACKEND_URL}/files/xml/export`}
            className="btn btn-primary btn-lg"
            target="_blank"
            rel="noreferrer"
            onClick={block}
          >
            <FontAwesomeIcon icon={'download'} /> Export .xml
          </a>
        </div>
      </form>
    </div>
  );
};

export default Export;
