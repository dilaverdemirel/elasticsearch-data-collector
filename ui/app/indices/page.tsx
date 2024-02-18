'use client'
import useSWR from 'swr'
import React from 'react';
import { Breadcrumb, Button, Form, Table } from 'react-bootstrap';
import { IoAddCircle } from 'react-icons/io5';
import { FaEdit } from 'react-icons/fa';
import { MdOutlineHistory } from 'react-icons/md';


function Indices() {
  const fetcher = (...args) => fetch(...args).then((res) => res.json())
  const { data, error } = useSWR('http://localhost:8080/indices', fetcher)

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
            {data.data.map((row) =>
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
              </tr>
            )}
          </tbody>
        </Table>
      </div>
    )
  }
  else {
    return <div>Loading...</div>
  }
}
export default Indices