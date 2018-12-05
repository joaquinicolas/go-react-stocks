import React from "react";
import { Table } from "reactstrap";

export function DataTable(props = { data: [] }) {
  return (
    <div style={{ marginTop: 10 }}>
      <Table responsive hover bordered>
        <thead>
          <tr>
            <th>#</th>
            <th>ID</th>
            <th>Age</th>
            <th>Name</th>
            <th>Email</th>
            <th>Phone</th>
          </tr>
        </thead>
        <tbody>
          {props.data.map((value, idx) => {
            return (
              <tr key={idx} onClick={() => props.onRowClicked(value.id)}>
                <th scope="row">{idx}</th>
                <td>{value.id}</td>
                <td>{value.age}</td>
                <td>{`${value.firstName} ${value.lastName}`}</td>
                <td>{value.email}</td>
                <td>{value.phone}</td>
              </tr>
            );
          })}
        </tbody>
      </Table>
    </div>
  );
}
