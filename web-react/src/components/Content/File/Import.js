import axios from 'axios';
import { BACKEND_URL } from 'constants/app';
import { useContext, useState } from 'react';
import ModalBox from 'components/UI/ModalBox';
import GetAxiosErrorMessage from 'components/Helpers/GetAxiosErrorMessage';
import ModalSpinner from 'components/UI/ModalSpinner';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import About from 'components/UI/About';
import UiContext from 'components/UI/Context';

const Import = () => {
  const { hasUserfile, setHasUserfile } = useContext(UiContext);
  const [modalShow, setModalShow] = useState(false);
  const [modalMessage, setModalMessage] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [userfile, setUserfile] = useState('');

  const onChange = (e) => {
    setUserfile(e.target.files[0]);
  };

  const onSubmit = async (e) => {
    e.preventDefault();
    if (userfile === '') {
      setModalMessage('Choose a file');
      setModalShow(true);
      return;
    }

    setIsLoading(true);

    const formData = new FormData();
    formData.append('userfile', userfile);
    try {
      console.log('hasUserfile:', hasUserfile);

      const r = await axios.post(`${BACKEND_URL}/files`, formData, {
        headers: {
          'content-type': 'multipart/form-data',
        },
      });

      if (r.status === 200) {
        setHasUserfile(true);
        setModalMessage(r.data);
        setModalShow(true);
      }
    } catch (err) {
      setModalMessage(GetAxiosErrorMessage(err));
      setModalShow(true);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div>
      <ModalSpinner show={isLoading} />
      <ModalBox show={modalShow} onHide={() => setModalShow(false)} message={modalMessage} />
      <h2>Import file</h2>
      <form className="row" onChange={onChange} onSubmit={onSubmit}>
        <div className="col-12">
          <label htmlFor="userfile" className="form-label">
            Select a <span className="text-black-50">.sav</span> file
          </label>
        </div>
        <div className="col-sm-12 col-md-10 col-lg-8">
          <input type="file" disabled={isLoading} className="form-control-file" id="userfile" name="userfile" />
        </div>
        <div className="col-12 mt-3">
          <button type="submit" disabled={isLoading} className="btn btn-primary btn-lg">
            <FontAwesomeIcon icon={'file'} /> Import
          </button>
        </div>
        <div className="col-12 mt-3">
          <p className="text-primary small">
            <b>Warning:</b> backup your .sav files before using this tool
          </p>
        </div>
      </form>
      <div className="fixed-bottom">
        <About />
      </div>
    </div>
  );
};

export default Import;
