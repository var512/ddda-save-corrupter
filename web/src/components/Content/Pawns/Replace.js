import axios from 'axios';
import { BACKEND_URL } from 'constants/app';
import { useState } from 'react';
import ModalBox from 'components/UI/ModalBox';
import GetAxiosErrorMessage from 'components/Helpers/GetAxiosErrorMessage';
import ModalSpinner from 'components/UI/ModalSpinner';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { useParams } from 'react-router-dom';

const Replace = () => {
  const { category } = useParams();
  const [modalShow, setModalShow] = useState(false);
  const [modalMessage, setModalMessage] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [pawnUserfile, setPawnUserfile] = useState('');

  const onChange = (e) => {
    setPawnUserfile(e.target.files[0]);
  };

  const onSubmit = async (e) => {
    e.preventDefault();
    if (pawnUserfile === '') {
      setModalMessage('Choose a file');
      setModalShow(true);
      return;
    }

    setIsLoading(true);

    const formData = new FormData();
    formData.append('pawnuserfile', pawnUserfile);
    formData.append('category', category);
    try {
      const r = await axios.post(`${BACKEND_URL}/pawns`, formData, {
        headers: {
          'content-type': 'multipart/form-data',
        },
      });

      if (r.status === 200) {
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
      <h2>Replace {category} pawn XML</h2>
      <form className="row" onChange={onChange} onSubmit={onSubmit}>
        <div className="col-12">
          <label htmlFor="pawnuserfile" className="form-label">
            Select a <span className="text-black-50">.xml</span> file (exported from this tool)
          </label>
        </div>
        <div className="col-sm-12 col-md-10 col-lg-8">
          <input type="file" disabled={isLoading} className="form-control-file" id="pawnuserfile" name="pawnuserfile" />
        </div>
        <div className="col-12 mt-3">
          <button type="submit" disabled={isLoading} className="btn btn-primary btn-lg">
            <FontAwesomeIcon icon={'people-arrows'} /> Replace
          </button>
        </div>
      </form>
    </div>
  );
};

export default Replace;
