import React, { Component } from "react";
import "./App.css";
import { Route, withRouter } from "react-router-dom";
import { Stock } from "./components/Stock";
import { SymbolForm } from "./components/Form";
import { Appbar } from "./components/Navbar";
import { Container, Row, Col, Modal, ModalHeader, ModalBody, ModalFooter, Button } from "reactstrap";
import { requestChart } from "./api";

class App extends Component {
  state = {
    data: [],
    modal: false,
    error: {
      title: '',
      message: ''
    }
  };

  constructor(props) {
    super(props);
    this.handleSubmit = this.handleSubmit.bind(this);
    this.toggle = this.toggle.bind(this);
  }

  toggle(err) {
    this.setState({
      error: err,
      modal: !this.state.modal
    });
  }
  handleSubmit(symbol1, symbol2, symbol3, range = '1m') {
    let queryString = "?symbols=";
    let dataTableHeaders = ['x'];
    const symbols = [symbol1, symbol2, symbol3];

    
    if (symbols.length === 0) {
      return this.toggle({
        title: "There is something wrong",
        message: "At least a symbol is expected to be provided"
      });
    }

    for (let i = 0; i < 3; i++) {
      if (symbols[i] && symbols[i].length > 1) {
        if (symbols[i]) {
          if (queryString.length > 9) {
            queryString += `,${symbols[i]}`;
          } else {
            queryString += `${symbols[i]}`;
          }
        }
        dataTableHeaders.push(symbols[i]);       
      }
    }
    queryString += `&range=${range}`;
    let self = this;
    requestChart(queryString)
    .then(values => this.parseData(dataTableHeaders, values))
    .catch(err =>  {
      if (err && err.error) {
        return self.toggle({
          title: "There is something wrong",
          message: err.error
        });
      }

      return self.toggle({
        title: "There is something wrong",
        message: 'Unexpected error'
      });
      
    });
  }

  
  parseData(dataTableHeaders, values) {
    let dataTable = [
      dataTableHeaders
    ];
    let i = 0; // column coordinate
    let j = 0; // row coordinate

    while (i < values[0].length) {
      let temp = [];
      while (j < values.length) {
        if (j === 0) {
          temp.push(values[j][i].date);
        }
        temp.push(values[j][i].close);
        j = j + 1;
      }
      i = i + 1;
      j = 0;

      dataTable.push(temp);
    }
    this.setState({
      data: dataTable
    }, () => this.props.history.push('/stocks'));
  }

  render() {
    return <Container fluid style={{ padding: 0 }}>
        <Row>
          <Col>
            <Appbar />
          </Col>
        </Row>
        <Container fluid>
          <Route path="/stocks" render={props => <Stock handleSubmit={this.handleSubmit} data={this.state.data} />} />
          <Route exact path="/" render={props => <SymbolForm handleSubmit={this.handleSubmit} />} />
        </Container>
        <Modal isOpen={this.state.modal} toggle={this.toggle}>
          <ModalHeader toggle={this.toggle}> {this.state.error.title} </ModalHeader>
          <ModalBody>
            {this.state.error.message}
          </ModalBody>
          <ModalFooter>
            <Button color="primary" onClick={this.toggle}>Get it</Button>
          </ModalFooter>
        </Modal>
      </Container>;
  }
}

export default withRouter(App);
