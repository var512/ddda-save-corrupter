import { useCallback, useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';

import axios from 'axios';
import { BACKEND_URL } from 'constants/app';
import GetAxiosErrorMessage from 'components/Helpers/GetAxiosErrorMessage';
import ModalSpinner from 'components/UI/ModalSpinner';
import ModalBox from 'components/UI/ModalBox';
import Information from 'components/Content/Pawns/Information';

const Pawn = () => {
  const { category } = useParams();
  const [modalShow, setModalShow] = useState(false);
  const [modalMessage, setModalMessage] = useState('');
  const [isLoading, setIsLoading] = useState(true);
  const [pawn, setPawn] = useState(null);

  const getPawns = useCallback(async (source) => {
    setPawn(null);

    try {
      const r = await axios.get(`${BACKEND_URL}/pawns/${category}`, {
        cancelToken: source.token,
      });

      if (r.status === 200) {
        setPawn(r.data.pawn);
      }
    } catch (err) {
      if (axios.isCancel(err)) {
        console.log(`cancelled request: ${source.token.reason}`);
        return;
      }

      setModalMessage(GetAxiosErrorMessage(err));
      setModalShow(true);
    } finally {
      setIsLoading(false);
    }
  }, [category]);

  useEffect(() => {
    const source = axios.CancelToken.source();

    getPawns(source);

    return () => {
      source.cancel();
    };
  }, [getPawns]);

  if (!pawn) {
    return <ModalSpinner show={true} />;
  }

  return (
    <div>
      <ModalSpinner show={isLoading} />
      <ModalBox show={modalShow} onHide={() => setModalShow(false)} message={modalMessage} />
      {
        pawn.Category !== ''
          ? <Information data={{
            category,
            pawn,
          }} />
          : <p className="text-black-50">
            Pawn data is empty.
          </p>
      }
    </div>
  );
};

export default Pawn;
