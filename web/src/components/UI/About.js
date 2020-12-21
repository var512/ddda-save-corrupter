import { useState, Fragment } from 'react';
import {
  Button, ListGroup, Toast,
} from 'react-bootstrap';

const About = () => {
  const [showAbout, setShowAbout] = useState(false);
  const toggleAbout = () => setShowAbout(!showAbout);

  return (
    <Fragment>
      <Toast onClose={toggleAbout} show={showAbout} animation={false} style={{ marginLeft: 40 }}>
        <Toast.Header>
          <strong className="mr-auto">Thanks</strong>
        </Toast.Header>
        <Toast.Body>
          <ListGroup variant="flush">
            <ListGroup.Item>fluffyquack.com</ListGroup.Item>
            <ListGroup.Item>github.com/Atvaark/ddda-save-editor</ListGroup.Item>
            <ListGroup.Item>github.com/Meem0/PawnManager</ListGroup.Item>
            <ListGroup.Item>github.com/beevik/etree</ListGroup.Item>
            <ListGroup.Item>github.com/gorilla/mux</ListGroup.Item>
            <ListGroup.Item>github.com/shurcooL/vfsgen</ListGroup.Item>
            <ListGroup.Item>github.com/goreleaser/goreleaser</ListGroup.Item>
            <ListGroup.Item>fontawesome.com</ListGroup.Item>
          </ListGroup>
        </Toast.Body>
      </Toast>
      <Button variant="link" onClick={toggleAbout}>
        ?
      </Button>
    </Fragment>
  );
};

export default About;
