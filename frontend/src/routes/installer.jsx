import { useState, useEffect } from "react"
import { Link } from "react-router-dom"
import Form from "react-bootstrap/Form"
import Button from "react-bootstrap/Button"
import Container from "react-bootstrap/Container"
import Row from "react-bootstrap/Row"
import Col from "react-bootstrap/Col"
import { useTranslation } from "react-i18next"

import "../i18n"
import { api } from "../global"

function Installer() {
  const { t } = useTranslation("Installer")
  const [path, setPath] = useState("")

  useEffect(() => {
    fetch(api + "installpath", {method: 'GET'})
      .then(data => data.text())
      .then(setPath)
      .catch(console.error)
  }, [])

  return (
    <Form>
      <Form.Group>
        <Form.Label>{t("Path")}</Form.Label>
        <Form.Control defaultValue={path}></Form.Control>
      </Form.Group>
      <Container className="mt-3">
        <Row className="justify-content-end">
          <Col xs="3">
            <Button className="w-100">
              {t("Install")}
            </Button>
          </Col>
          <Col xs="3">
            <Button className="w-100" onClick={ () => fetch()}>
              {t("Exit")}
            </Button>
          </Col>
        </Row>
      </Container>
    </Form>
    // <Form>
    //   <Form.Group>
    //     <Form.Label>Install path</Form.Label>
    //     <Form.Control />
    //   </Form.Group>
    //   {/* <Button variant="primary">Install</Button> */}
    //     <Link to="/www/addserver">Click Here</Link>
    // </Form>
  )
}

export default Installer
