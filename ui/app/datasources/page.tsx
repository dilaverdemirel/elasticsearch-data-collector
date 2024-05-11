'use client'
import useSWR from 'swr'
import { Breadcrumb, Button, Form, Table } from 'react-bootstrap';
import { IoAddCircle } from 'react-icons/io5';
import { FaEdit } from 'react-icons/fa';

const fetcher = (...args: any[]) => fetch(...args).then((res) => res.json())

function Datasources() {
  const { data, error } = useSWR('http://localhost:8080/datasources', fetcher)

  if (error) return <div>Failed to load {error}</div>
  if (!data) return <div>Loading...</div>

  const columns = [
    {
      key: "Name",
      label: "Name",
    },
    {
      key: "UserName",
      label: "User Name",
    },
    {
      key: "DriverName",
      label: "Driver Name",
    },
    {
      key: "MaxPoolSize",
      label: "Max Pool Size",
    },
    {
      key: "MinIdle",
      label: "Min Idle",
    },
  ];

  return (
    <div style={{ width: "100%" }}>

      <Breadcrumb>
        <Breadcrumb.Item href="/">Home</Breadcrumb.Item>
        <Breadcrumb.Item active>Datasources</Breadcrumb.Item>
      </Breadcrumb>


      <div style={{ width: "100%", paddingTop: "5px", paddingBottom: "5px" }}>
        <Button href='/datasources/form'><div style={{ display: 'flex' }}><IoAddCircle size={25} /> Add New</div></Button>
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
            <tr key={row.ID}>
              {columns.map((column, index) =>
                <td key={row.ID + "-" + index}>
                  {column.key !== 'Valid' && column.key !== 'Scheduled' ? row[column.key] :
                    <Form.Check type="checkbox" defaultChecked={row[column.key]} disabled={true} />
                  }
                </td>
              )}
              <td>
                <Button variant='link' href={'/datasources/form/' + row.ID}><FaEdit size={20} /></Button>
              </td>
            </tr>
          )}
        </tbody>
      </Table>
    </div>
  )
}

export default Datasources