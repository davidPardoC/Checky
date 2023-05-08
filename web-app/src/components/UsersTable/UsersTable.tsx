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
import useUsersStore from "@/store/users.store";

const UsersTable = () => {
  const [showAddModal, setShowAddModal] = useState(false);

  const users = useUsersStore((state) => state.users);

  const openAddUserModal = () => {
    setShowAddModal(true);
  };

  const closeModal = () => {
    setShowAddModal(false);
  };

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
              {users.map((user) => (
                <Tr key={user.id}>
                  <Td>{user.name}</Td>
                  <Td>{user.lastname}</Td>
                  <Td>{user.email}</Td>
                  <Td>{user.role}</Td>
                </Tr>
              ))}
            </Tbody>
          </Table>
        </TableContainer>
      </Container>
    </>
  );
};

export default UsersTable;
