'use client'
import { Breadcrumb, Button, Card, Col, Form, Modal, Row } from "react-bootstrap"
import { useEffect, useState } from "react";
import { useParams, useRouter } from 'next/navigation'
import { toast } from "react-toastify";
import Datasource from "../../datasource";
import { createDatasource, deleteDatasource, getDatasourceById, updateDatasource } from "../../service";

export default function DatasourceForm() {

  const router = useRouter()
  const params = useParams<{ id: string; }>()
  const [validated, setValidated] = useState(false);
  const [datasource, setDatasource] = useState(new Datasource());
  const [showDeleteForm, setShowDeleteForm] = useState(false);


  const loadDatasource = (id: string) => {
    getDatasourceById(id)
      .then(res => {
        setDatasource(res.data)
      })
  }

  useEffect(() => {
    if (params?.id?.length > 0) {
      loadDatasource(params.id[0])
    }
  }, [])

  const redirectToDatasourcesList = () => {
    setTimeout(() => {
      router.push('/datasources')
    }, 2000);
  }

  const handleSubmit = (event) => {
    const form = event.currentTarget;

    event.preventDefault();
    if (form.checkValidity() === false) {
      event.preventDefault();
      event.stopPropagation();
    } else {
      if (datasource.ID) {
        console.log("Updating.......")
        updateDatasource(datasource)
          .then((res) => {
            console.log(res.status)
            console.log(res.data)
            toast.success("Successfully updated.");
            redirectToDatasourcesList()
          })
          .catch((err) => {
            toast.error("" + err);
            console.error(err)
          })
      } else {
        console.log("Creating.......")
        createDatasource(datasource)
          .then((res) => {
            console.log(res.status)
            console.log(res.data)
            toast.success("Successfully created.");
            redirectToDatasourcesList()
          })
          .catch((err) => {
            toast.error("" + err);
            console.error(err)
          })
      }
    }

    setValidated(true);
  };


  const deleteDatasourceFnc = () => {
    deleteDatasource(datasource.ID)
      .then((res) => {
        console.log(res.status)
        console.log(res.data)
        toast.success("Successfully deleted.");
        redirectToDatasourcesList()
      })
      .catch((err) => {
        toast.error("" + err);
        console.error(err)
      })
  }

  return (
    <div style={{ width: "100%" }}>
      <Breadcrumb>
        <Breadcrumb.Item href="/">Home</Breadcrumb.Item>
        <Breadcrumb.Item href="/datasources">Datasources</Breadcrumb.Item>
        <Breadcrumb.Item active>Edit Datasource</Breadcrumb.Item>
      </Breadcrumb>

      <Form noValidate validated={validated} onSubmit={handleSubmit} style={{ width: "100%" }}>
        <Row className="mb-3">
          <Form.Group as={Col} md="12">
            <Form.Label>Datasource Name</Form.Label>
            <Form.Control
              required
              type="text"
              placeholder="demo-ds"
              value={datasource.Name}
              onChange={(e) => {
                setDatasource({ ...datasource, Name: e.target.value })
              }}
            />
          </Form.Group>
        </Row>
        <Row className="mb-3">
          <Form.Group as={Col} md="12">
            <Form.Label>Connection String</Form.Label>
            <Form.Control
              required
              type="text"
              placeholder="tcp(127.0.0.1:3306)/es-data-collector?parseTime=true"
              value={datasource.ConnectionString}
              onChange={(e) => {
                setDatasource({ ...datasource, ConnectionString: e.target.value })
              }}
            />
          </Form.Group>
        </Row>
        <Row className="mb-3">
          <Form.Group as={Col} className="mb-3" md="6">
            <Form.Label>Max Pool Size</Form.Label>
            <Form.Control
              required
              type="text"
              placeholder="10"
              value={datasource.MaxPoolSize}
              onChange={(e) => {
                setDatasource({ ...datasource, MaxPoolSize: Number(e.target.value) })
              }}
            />
          </Form.Group>
          <Form.Group as={Col} className="mb-3" md="6">
            <Form.Label>Min Idle</Form.Label>
            <Form.Control
              required
              type="text"
              placeholder="2"
              value={datasource.MinIdle}
              onChange={(e) => {
                setDatasource({ ...datasource, MinIdle: Number(e.target.value) })
              }}
            />
          </Form.Group>
        </Row>
        <Row className="mb-3">
          <Form.Group as={Col} className="mb-3" md="6">
            <Form.Label>Username</Form.Label>
            <Form.Control
              required
              type="text"
              placeholder="user"
              value={datasource.UserName}
              onChange={(e) => {
                setDatasource({ ...datasource, UserName: e.target.value })
              }}
            />
          </Form.Group>
          <Form.Group as={Col} className="mb-3" md="6">
            <Form.Label>Password</Form.Label>
            <Form.Control
              required
              type="password"
              placeholder="pwd"
              value={datasource.DbPassword}
              onChange={(e) => {
                setDatasource({ ...datasource, DbPassword: e.target.value })
              }}
            />
          </Form.Group>
        </Row>
        <Row className="mb-3">
          <Form.Group as={Col} className="mb-3">
            <Form.Label>Driver Name</Form.Label>
            <Form.Control
              required
              type="text"
              placeholder="mysql"
              value={datasource.DriverName}
              onChange={(e) => {
                setDatasource({ ...datasource, DriverName: e.target.value })
              }}
            />
          </Form.Group>
        </Row>

        <Row className="mb-3">
          <Form.Group as={Col} md="2" className="mb-3">
            <Button type="submit">Save Changes</Button>
          </Form.Group>
          <Form.Group as={Col} md="10" className="mb-3" style={{ visibility: datasource.ID ? 'visible' : 'hidden' }}>
            <Button variant="danger" onClick={() => setShowDeleteForm(true)}>Delete Datasource</Button>
          </Form.Group>
        </Row>
      </Form>

      {/* Delete Modal */}
      <Modal show={showDeleteForm} aria-labelledby="contained-modal-title-vcenter" centered>
        <Modal.Header>
          <Modal.Title>Delete Datasource</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <p>Do you want to continue? Datasource will be deleted!</p>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={() => setShowDeleteForm(false)}>
            Close
          </Button>
          <Button variant="danger" onClick={() => { deleteDatasourceFnc(); setShowDeleteForm(false); }}>
            Delete
          </Button>
        </Modal.Footer>
      </Modal>
    </div>
  )

}