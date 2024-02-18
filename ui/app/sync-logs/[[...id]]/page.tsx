'use client'
import useSWR from 'swr'
import React, { useEffect, useState } from 'react';
import { Breadcrumb, Form, Table } from 'react-bootstrap';
import { findByIndexId } from '../service';
import { useParams } from 'next/navigation';


function SyncLogs() {
  const params = useParams<{ id: string; }>()

  const [data, setData] = useState(null);


  const loadLogs = (id: string) => {
    findByIndexId(id)
      .then(res => {
        setData(res.data)
      })
  }

  useEffect(() => {

    if (params?.id?.length > 0) {
      loadLogs(params.id[0])
    }

  }, [])

  const columns = [
    {
      key: "DocumentCount",
      label: "Transfered Document Count",
    },
    {
      key: "StartDate",
      label: "Start Date",
    },
    {
      key: "EndDate",
      label: "End Date",
    },
    {
      key: "ExecutionDuration",
      label: "Execution Duration",
    },
    {
      key: "Status",
      label: "Status",
    },
    {
      key: "StatusMessage",
      label: "Status Message",
    },
  ];

  if (data) {
    return (
      <div style={{ width: "100%" }}>
        <Breadcrumb>
          <Breadcrumb.Item href="/">Home</Breadcrumb.Item>
          <Breadcrumb.Item href='/indices'>Indices</Breadcrumb.Item>
          <Breadcrumb.Item active>Sync Logs</Breadcrumb.Item>
        </Breadcrumb>

        <Table striped bordered hover>
          <thead>
            <tr>
              {columns.map((column) =>
                <th key={column.key}>{column.label}</th>
              )}
            </tr>
          </thead>
          <tbody>
            {data.map((row) =>
              <tr key={row.ID + "tr"}>
                {columns.map((column) =>
                  <td key={row.ID + "td1" + column.key}>
                    {row[column.key]}
                  </td>
                )}
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
export default SyncLogs