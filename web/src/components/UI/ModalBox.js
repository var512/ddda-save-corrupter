import { Button, Modal } from 'react-bootstrap';

const ModalBox = (props) => (
  <Modal size="lg" animation="false" centered="true" {...props}>
    <Modal.Body>
      <h4 className="m-0 p-0">{props.message}</h4>
    </Modal.Body>
    <Modal.Footer>
      <Button size="lg" onClick={props.onHide}>Close</Button>
    </Modal.Footer>
  </Modal>
);

export default ModalBox;
