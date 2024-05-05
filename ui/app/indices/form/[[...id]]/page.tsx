'use client'
import { Breadcrumb, Button, Card, Col, Form, Modal, OverlayTrigger, Row, Table, Tooltip } from "react-bootstrap"
import { useEffect, useState } from "react";
import { createIndex, deleteIndex, getIndexById, scheduleIndexDataSync, startDataSyncImmediately, unscheduleIndexDataSync, updateIndex } from "../../service";
import Index, { ScheduleIndex } from "../..";
import { useParams, useRouter } from 'next/navigation'
import { toast } from "react-toastify";
import { getDatasources, getQueryPreviewData } from "@/app/datasources/service";
import { QueryPreviewDataDTO, QueryPreviewResultDTO } from "@/app/datasources/datasource";
import { MdOutlineHistory, MdPreview, MdSchedule } from "react-icons/md";
import { VscOpenPreview } from "react-icons/vsc";

export default function IndexForm() {

  const router = useRouter()
  const params = useParams<{ id: string; }>()
  const [validated, setValidated] = useState(false);
  const [validatedForScheduleForm, setValidatedForScheduleForm] = useState(false);
  const [dsdata, setDsdata] = useState(null);
  const [index, setIndex] = useState(new Index());
  const [scheduleIndex, setScheduleIndex] = useState(new ScheduleIndex());
  const [showScheduleForm, setShowScheduleForm] = useState(false);
  const [showUnscheduleForm, setShowUnscheduleForm] = useState(false);
  const [showDeleteForm, setShowDeleteForm] = useState(false);
  const [showPreviewResult, setShowPreviewResult] = useState(false);
  const [queryPreviewResult, setQueryPreviewResult] = useState(new QueryPreviewResultDTO());


  const loadIndex = (id: string) => {
    getIndexById(id)
      .then(res => {
        setIndex(res.data)
        setScheduleIndex({
          ...scheduleIndex,
          CronExpression: res.data.CronExpression,
          DocumentIdField: res.data.DocumentField,
          SyncType: res.data.SyncType,
          IndexId: res.data.ID
        })
      })
  }

  const sqlPreview = () => {
    const previewMetadata = new QueryPreviewDataDTO()
    previewMetadata.DataSourceId = index.DataSourceId
    previewMetadata.Query = index.SqlQuery
    getQueryPreviewData(previewMetadata)
      .then(res => {
        setQueryPreviewResult(res.data)
      })
  }

  useEffect(() => {
    getDatasources().then(res => {
      setDsdata(res.data)
    })

    if (params?.id?.length > 0) {
      loadIndex(params.id[0])
    }

  }, [])

  const redirectToIndicesList = () => {
    setTimeout(() => {
      router.push('/indices')
    }, 2000);
  }

  const handleSubmit = (event) => {
    const form = event.currentTarget;

    event.preventDefault();
    if (form.checkValidity() === false) {
      event.preventDefault();
      event.stopPropagation();
    } else {
      if (index.ID) {
        console.log("Updating.......")
        updateIndex(index)
          .then((res) => {
            console.log(res.status)
            console.log(res.data)
            toast.success("Successfully updated.");
            redirectToIndicesList()
          })
          .catch((err) => {
            toast.error("" + err);
            console.error(err)
          })
      } else {
        console.log("Creating.......")
        createIndex(index)
          .then((res) => {
            console.log(res.status)
            console.log(res.data)
            toast.success("Successfully created.");
            redirectToIndicesList()
          })
          .catch((err) => {
            toast.error("" + err);
            console.error(err)
          })
      }
    }

    setValidated(true);
  };

  const handleSubmitForScheduleForm = (event) => {
    const form = event.currentTarget;
    event.preventDefault();

    if (form.checkValidity() === false) {
      event.preventDefault();
      event.stopPropagation();
    } else {
      console.log("Scheduling.......")
      scheduleIndexDataSync(scheduleIndex)
        .then((res) => {
          console.log(res.status)
          console.log(res.data)
          toast.success("Successfully scheduled.");
          loadIndex(index.ID)
          setShowScheduleForm(false)
        })
        .catch((err) => {
          toast.error("" + err);
          console.error(err)
        })
    }
    setValidatedForScheduleForm(true);
  };

  const unschedule = () => {
    unscheduleIndexDataSync(index.ID)
      .then((res) => {
        console.log(res.status)
        console.log(res.data)
        loadIndex(index.ID)
        toast.success("Successfully unscheduled.");
      })
      .catch((err) => {
        toast.error("" + err);
        console.error(err)
      })
  }

  const deleteIndexFnc = () => {
    deleteIndex(index.ID)
      .then((res) => {
        console.log(res.status)
        console.log(res.data)
        toast.success("Successfully deleted.");
        redirectToIndicesList()
      })
      .catch((err) => {
        toast.error("" + err);
        console.error(err)
      })
  }

  const renderSqlQueryPreviewTooltip = (props) => (
    <Tooltip id="button-tooltip" {...props}>
      SQL Query Preview Results
    </Tooltip>
  );

  return (
    <div style={{ width: "100%" }}>
      <Breadcrumb>
        <Breadcrumb.Item href="/">Home</Breadcrumb.Item>
        <Breadcrumb.Item href="/indices">Indices</Breadcrumb.Item>
        <Breadcrumb.Item active>Edit Index</Breadcrumb.Item>
      </Breadcrumb>

      <Form noValidate validated={validated} onSubmit={handleSubmit} style={{ width: "100%" }}>
        <Row className="mb-3">
          <Form.Group as={Col} md="12">
            <Form.Label>Index Name</Form.Label>
            <Form.Control
              required
              type="text"
              placeholder="demo-index"
              value={index.Name}
              onChange={(e) => {
                setIndex({ ...index, Name: e.target.value })
              }}
            />
          </Form.Group>
        </Row>
        <Row className="mb-3">
          <Form.Group as={Col} md="12">
            <Form.Label>Description</Form.Label>
            <Form.Control
              required
              type="text"
              value={index.Description}
              onChange={(e) => {
                setIndex({ ...index, Description: e.target.value })
              }}
            />
          </Form.Group>
        </Row>
        <Row className="mb-3">
          <Form.Group as={Col} className="mb-3">
            <Form.Label>SQL Query</Form.Label>
            <Form.Control required as="textarea" rows={3} placeholder="select * from customers"
              onChange={(e) => setIndex({ ...index, SqlQuery: e.target.value })}
              value={index.SqlQuery} />
          </Form.Group>
        </Row>
        <Row className="mb-3">
          <Form.Group as={Col} md="5" className="mb-3">
            <Form.Label>Datasource</Form.Label>
            <Form.Select aria-label="Datasource" required onChange={(e) => setIndex({ ...index, DataSourceId: e.target.value })}
              value={index.DataSourceId}>
              <option></option>
              {dsdata && dsdata.data.map((ds) => {
                return <option key={ds["ID"]} value={ds["ID"]}>{ds["Name"]}</option>
              })}
            </Form.Select>
          </Form.Group>

          <Form.Group as={Col} md="1" className="mb-3">

            <Button variant="link"
              style={{ marginTop: 30 }}
              onClick={() => { setShowPreviewResult(true); sqlPreview(); }}>
              <VscOpenPreview size={30} />
            </Button>

          </Form.Group>

          <Form.Group as={Col} md="2" className="mb-3">
            <Form.Check
              type="switch"
              id="valid"
              label="Valid"
              checked={index.Valid}
              style={{ marginTop: 40 }}
              onChange={(e) => setIndex({ ...index, Valid: e.target.checked })}
            />
          </Form.Group>
          <Form.Group as={Col} md="4" className="mb-3">
            <Button variant="link" onClick={() => setShowScheduleForm(true)}
              style={{ display: index.ID && !index.Scheduled ? 'block' : 'none', marginTop: 35 }}>
              Schedule Data Sync
            </Button>
            <Card style={{ display: index.ID && index.Scheduled ? 'block' : 'none', width: '100%' }}>
              <Card.Body>
                <Card.Title>Scheduling Details</Card.Title>
                <Card.Text>
                  Index data sync was scheduled at <b>{index.CronExpression}</b> with
                  <b>{index.SyncType == 'RELOAD_ALL' ? ' Reload All' : ' Iterative'}</b> syncronization
                  type. Elasticsearch document id field name is <b>{index.DocumentField}</b> on data sync.
                </Card.Text>
                <Card.Link href="#" onClick={() => setShowUnscheduleForm(true)}
                  style={{ visibility: index.ID && index.Scheduled ? 'visible' : 'hidden' }}>Unschedule Data Sync</Card.Link>
              </Card.Body>
            </Card>
          </Form.Group>
        </Row>
        <Row className="mb-3">
          <Form.Group as={Col} md="2" className="mb-3">
            <Button type="submit">Save Changes</Button>
          </Form.Group>
          <Form.Group as={Col} md="10" className="mb-3" style={{ visibility: index.ID ? 'visible' : 'hidden' }}>
            <Button variant="danger" onClick={() => setShowDeleteForm(true)}>Delete Index</Button>
          </Form.Group>
        </Row>
      </Form>

      {/* Schedule Modal */}
      <Modal show={showScheduleForm} aria-labelledby="contained-modal-title-vcenter" centered>
        <Form noValidate validated={validatedForScheduleForm} onSubmit={handleSubmitForScheduleForm} style={{ width: "100%" }}>
          <Modal.Header>
            <Modal.Title>Schedule Data Sync</Modal.Title>
          </Modal.Header>
          <Modal.Body>

            <Form.Group className="mb-3" controlId="exampleForm.ControlInput1">
              <Form.Label>Cron Expression</Form.Label>
              <Form.Control
                type="text"
                placeholder="0/15 0 * * ?"
                onChange={(e) => setScheduleIndex({ ...scheduleIndex, CronExpression: e.target.value })}
                value={scheduleIndex.CronExpression}
              />
            </Form.Group>
            <Form.Group
              className="mb-3"
            >
              <Form.Label>Document Id Field</Form.Label>
              <Form.Control
                type="text"
                placeholder="id"
                required
                onChange={(e) => setScheduleIndex({ ...scheduleIndex, DocumentIdField: e.target.value })}
                value={scheduleIndex.DocumentIdField}
              />
            </Form.Group>
            <Form.Group
              className="mb-3"
            >
              <Form.Label>Sync Type</Form.Label>
              <Form.Select aria-label="Sync Type" required
                onChange={(e) => setScheduleIndex({ ...scheduleIndex, SyncType: e.target.value })}
                value={scheduleIndex.SyncType}
              >
                <option></option>
                <option value="ITERATIVE">Iterative</option>
                <option value="RELOAD_ALL">Reload All</option>
              </Form.Select>
            </Form.Group>

          </Modal.Body>
          <Modal.Footer>
            <Button variant="secondary" onClick={() => setShowScheduleForm(false)}>
              Close
            </Button>
            <Button type="submit" variant="primary" >
              Schedule
            </Button>
          </Modal.Footer>
        </Form>
      </Modal>

      {/* Unschedule Modal */}
      <Modal show={showUnscheduleForm} aria-labelledby="contained-modal-title-vcenter" centered>
        <Modal.Header>
          <Modal.Title>Unschedule Data Sync</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <p>Do you want to continue? Index data sync will be unscheduled!</p>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={() => setShowUnscheduleForm(false)}>
            Close
          </Button>
          <Button variant="primary" onClick={() => { unschedule(); setShowUnscheduleForm(false); }}>
            Unschedule
          </Button>
        </Modal.Footer>
      </Modal>

      {/* Delete Modal */}
      <Modal show={showDeleteForm} aria-labelledby="contained-modal-title-vcenter" centered>
        <Modal.Header>
          <Modal.Title>Delete Index</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <p>Do you want to continue? Index will be deleted!</p>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={() => setShowDeleteForm(false)}>
            Close
          </Button>
          <Button variant="danger" onClick={() => { deleteIndexFnc(); setShowDeleteForm(false); }}>
            Delete
          </Button>
        </Modal.Footer>
      </Modal>

      {/* Sql Query Preview Result Modal */}
      <Modal fullscreen show={showPreviewResult} aria-labelledby="contained-modal-title-vcenter">
        <Modal.Header>
          <Modal.Title>SQL Query Preview</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Table striped bordered hover>
            <thead>
              <tr key={"header"}>
                {queryPreviewResult.MetaDataList.map((column) =>
                  <th key={column.FieldName}>{column.FieldName}</th>
                )}
              </tr>
            </thead>
            <tbody>
              {queryPreviewResult.ExampleData.map((row, rowindex) =>
                <tr key={rowindex + "tr"}>
                  {queryPreviewResult.MetaDataList.map((column, colindex) =>
                    <td key={colindex + "td1" + column.FieldName}>
                      {row[column.FieldName]}
                    </td>
                  )}
                </tr>
              )}
            </tbody>
          </Table>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={() => setShowPreviewResult(false)}>
            Close
          </Button>
        </Modal.Footer>
      </Modal>
    </div>
  )

}