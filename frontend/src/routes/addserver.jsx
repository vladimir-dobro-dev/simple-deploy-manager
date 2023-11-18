import Form from "react-bootstrap/Form"
import Button from "react-bootstrap/Button"
import { Link } from 'react-router-dom'

function AddServer() {
  return (
    <Form>
      <Form.Group>
        <Form.Label>IP Address</Form.Label>
        <Form.Control />
      </Form.Group>
      <Form.Group>
        <Form.Label>User name</Form.Label>
        <Form.Control />
      </Form.Group>
      {/* <Button variant="primary">Install</Button> */}
      <Link to="/www/installer">Click Here</Link>
    </Form>
  )
}

export default AddServer
