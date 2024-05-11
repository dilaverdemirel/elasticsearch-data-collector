'use client'
import useSWR from 'swr'
import React, { useState } from 'react';
import { Breadcrumb, Button, Form, Modal, Table } from 'react-bootstrap';
import { IoAddCircle } from 'react-icons/io5';
import { FaEdit } from 'react-icons/fa';
import { MdOutlineHistory, MdSyncAlt } from 'react-icons/md';
import { startDataSyncImmediately } from './service';
import { toast } from 'react-toastify';


function Indices() {
  const fetcher = (...args: any[]) => fetch(...args).then((res) => res.json())
  const { data, error } = useSWR('http://localhost:8080/indices', fetcher)

  const [showOneTimeStartSync, setShowOneTimeStartSync] = useState(false);
  const [selectedIndexId, setSelectedIndexId] = useState("");

  if (error) return <div>Failed to load {error}</div>
  //if (!data) return <div>Loading...</div>

  const columns = [
    {
      key: "Name",
      label: "Name",
    },
    {
      key: "Description",
      label: "Description",
    },
    {
      key: "SyncType",
      label: "Sync Type",
    },
    {
      key: "Valid",
      label: "Valid",
    },
    {
      key: "DocumentField",
      label: "Document Field",
    },
    {
      key: "Scheduled",
      label: "Scheduled",
    },
    {
      key: "CronExpression",
      label: "Cron Expression",
    },
  ];

  const startDataSync = () => {
    startDataSyncImmediately(selectedIndexId)
      .then(res => {
        toast.success("Successfully started.");
      }).catch((err) => {
        toast.error("" + err);
        console.error(err)
      })
  }

  if (data) {
    return (
      <div style={{ width: "100%" }}>
        <Breadcrumb>
          <Breadcrumb.Item href="/">Home</Breadcrumb.Item>
          <Breadcrumb.Item active>Indices</Breadcrumb.Item>
        </Breadcrumb>

        <div style={{ width: "100%", paddingTop: "5px", paddingBottom: "5px" }}>
          <Button href='/indices/form'><div style={{ display: 'flex' }}><IoAddCircle size={25} /> Add New</div></Button>
        </div>

        <Table striped bordered hover>
          <thead>
            <tr>
              {columns.map((column) =>
                <th key={column.key}>{column.label}</th>
              )}
            </tr>
          </thead>
          <tbody>
            {data.data.map((row : any) =>
              <tr key={row.ID + "tr"}>
                {columns.map((column) =>
                  <td key={row.ID + "td1" + column.key}>
                    {column.key !== 'Valid' && column.key !== 'Scheduled' ? row[column.key] :
                      <Form.Check key={row.ID + "check" + column.key} type="checkbox" defaultChecked={row[column.key]} disabled={true} />
                    }
                  </td>
                )}
                <td key={row.ID + "td2"}>
                  <Button key={row.ID + "edit"} variant='link' href={'/indices/form/' + row.ID}><FaEdit size={20} /></Button>
                </td>
                <td key={row.ID + "td3"}>
                  <Button key={row.ID + "history"} variant='link' href={'/sync-logs/' + row.ID}><MdOutlineHistory size={20} /></Button>
                </td>
                <td key={row.ID + "td4"}>
                  <Button variant="link" onClick={() => { setSelectedIndexId(row.ID); setShowOneTimeStartSync(true); }}>
                    <MdSyncAlt size={20} />
                  </Button>
                </td>
              </tr>
            )}
          </tbody>
        </Table>

        {/* OneTime Sync Modal */}
        <Modal show={showOneTimeStartSync} aria-labelledby="contained-modal-title-vcenter" centered>
          <Modal.Header>
            <Modal.Title>Start OneTime Data Sync</Modal.Title>
          </Modal.Header>
          <Modal.Body>
            <p>Do you want to continue?</p>
          </Modal.Body>
          <Modal.Footer>
            <Button variant="secondary" onClick={() => setShowOneTimeStartSync(false)}>
              Close
            </Button>
            <Button variant="primary" onClick={() => { startDataSync(); setShowOneTimeStartSync(false); }}>
              Start data sync...
            </Button>
          </Modal.Footer>
        </Modal>
      </div>
    )
  }
  else {
    return <div>Loading...</div>
  }
}
export default Indices