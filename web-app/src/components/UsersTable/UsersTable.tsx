import { Button, Container } from "@chakra-ui/react";
import React, { useState } from "react";
import {
  Table,
  Thead,
  Tbody,
  Tfoot,
  Tr,
  Th,
  Td,
  TableContainer,
} from "@chakra-ui/react";
import { CreateUserForm } from "../SignupForm/CreateUserForm";

const UsersTable = () => {
  const [showAddModal, setShowAddModal] = useState(false);

  const openAddUserModal = () => {
    setShowAddModal(true)
  }

  const closeModal = () => {
    setShowAddModal(false)
  }

  return (
    <>
      {showAddModal && <CreateUserForm onClose={closeModal} />}
      <Container paddingTop={5}>
        <Button bg={"primary"} color={"white"} onClick={openAddUserModal}>
          New user
        </Button>
        <TableContainer>
          <Table variant="simple">
            <Thead>
              <Tr>
                <Th>Name</Th>
                <Th>Last Name</Th>
                <Th>Email</Th>
                <Th>Actions</Th>
              </Tr>
            </Thead>
            <Tbody>
              <Tr>
                <Td>David</Td>
                <Td>Pardo</Td>
                <Td>pardodavid10@gmail.com</Td>
                <Td>Actions</Td>
              </Tr>
            </Tbody>
          </Table>
        </TableContainer>
      </Container>
    </>
  );
};

export default UsersTable;
