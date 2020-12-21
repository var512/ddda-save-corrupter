import { Modal } from 'react-bootstrap';

const ModalSpinner = (props) => (
  <Modal size="sm" backdrop="static" keyboard={false} animation="false" centered="true" {...props}>
    <Modal.Body>
      <div className="text-center pt-4 pb-4">
        <div className="spinner-border text-black-50" role="status">
        </div>
        <h5 className="m-0 p-0 mt-3 text-black-50">Loading...</h5>
      </div>
    </Modal.Body>
  </Modal>
);

export default ModalSpinner;
