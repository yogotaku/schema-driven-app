import { useState, type FC } from 'react';
import './App.css';
import {
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
} from '@mui/material';
import {
  getUserById,
  createUser as createUserFunc,
  patchUser as patchUserFunc,
  getPets,
} from 'api';
import { type User, type NewUser, type Pet } from 'type';
import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import Grid from '@mui/material/Grid';
import TextField from '@mui/material/TextField';
import Typography from '@mui/material/Typography';

const initialUser: User = {
  id: 0,
  firstName: '',
  lastName: '',
  email: '',
  emailVerified: false,
  dateOfBirth: '',
  createDate: '',
};

const initialNewUser: NewUser = {
  firstName: '',
  lastName: '',
  email: '',
  dateOfBirth: '',
};

const App: FC = () => {
  const [id, setId] = useState<number>(0);
  const [patchUser, setPatchUser] = useState<User>(initialUser);
  const [createUser, setCreateUser] = useState<NewUser>(initialNewUser);
  const [tags, setTags] = useState<string[]>([]);
  const [limit, setLimit] = useState<number | null>(null);
  const [pets, setPets] = useState<Pet[]>([]);

  const handleOnClickUserSearch = async () => {
    const user = await getUserById(id);
    if (user !== null) {
      setPatchUser({ ...user });
    }
  };

  const handleOnClickCreate = async () => {
    const user = await createUserFunc(createUser);
    if (user !== null) {
      window.alert(`id=${user.id}のユーザーを登録しました`);
    } else {
      window.alert('ユーザーの登録に失敗しました');
    }
  };

  const handleOnClickUpdate = async () => {
    const res = await patchUserFunc(id, patchUser);
    if (res) {
      window.alert('ユーザー情報の更新に成功しました');
    } else {
      window.alert('ユーザー情報の更新に失敗しました');
    }
  };

  const handleOnClickPetSearch = async () => {
    const pets = await getPets(tags, limit);
    if (pets !== null) {
      setPets(pets);
    }
  };

  return (
    <Box
      sx={{
        flexGrow: 1,
        width: '100%',
        height: '100vh',
      }}
    >
      <Grid
        container
        spacing={2}
        justifyContent="space-around"
        sx={{ backgroundColor: 'lightcyan' }}
      >
        {/* ユーザー検索・更新ここから */}
        <Grid
          item
          xs={5}
          sx={{
            backgroundColor: 'azure',
            minHeight: '450px',
          }}
        >
          <Typography variant="h5" textAlign={'left'}>
            ユーザー検索/更新
          </Typography>
          <Grid container sx={{ paddingTop: '1rem' }} justifyContent={'left'}>
            <Grid item xs={8} sx={{ marginBottom: '1rem' }}>
              <TextField
                label="UserID"
                variant="outlined"
                sx={{ width: '100%' }}
                onChange={(e) => {
                  const id = parseInt(e.target.value, 10);
                  if (!isNaN(id)) {
                    setId(id);
                  } else {
                    setId(0);
                  }
                }}
                value={id !== 0 ? id : ''}
              ></TextField>
            </Grid>
            <Grid item xs={4} sx={{ marginBottom: '1rem' }}>
              <Button
                variant="outlined"
                sx={{ height: '100%' }}
                onClick={async () => {
                  await handleOnClickUserSearch();
                }}
              >
                Search
              </Button>
            </Grid>
            <Grid item xs={8} sx={{ marginBottom: '1rem' }}>
              <TextField
                label="FirstName"
                variant="outlined"
                sx={{ width: '100%' }}
                onChange={(e) => {
                  setPatchUser((prev) => {
                    const newUser = { ...prev };
                    newUser.firstName = e.target.value;

                    return newUser;
                  });
                }}
                value={patchUser.firstName}
              ></TextField>
            </Grid>
            <Grid item xs={4} sx={{ marginBottom: '1rem' }}>
              <Button
                variant="outlined"
                sx={{ height: '100%' }}
                onClick={async () => {
                  await handleOnClickUpdate();
                }}
              >
                Update
              </Button>
            </Grid>
            <Grid item xs={8} sx={{ marginBottom: '1rem' }}>
              <TextField
                label="LastName"
                variant="outlined"
                sx={{ width: '100%' }}
                onChange={(e) => {
                  setPatchUser((prev) => {
                    const newUser = { ...prev };
                    newUser.lastName = e.target.value;

                    return newUser;
                  });
                }}
                value={patchUser.lastName}
              ></TextField>
            </Grid>
            <Grid item xs={8} sx={{ marginBottom: '1rem' }}>
              <TextField
                label="Email"
                variant="outlined"
                sx={{ width: '100%' }}
                onChange={(e) => {
                  setPatchUser((prev) => {
                    const newUser = { ...prev };
                    newUser.email = e.target.value;

                    return newUser;
                  });
                }}
                value={patchUser.email}
              ></TextField>
            </Grid>
            <Grid item xs={8} sx={{ marginBottom: '1rem' }}>
              <TextField
                label="DateOfBirth"
                variant="outlined"
                sx={{ width: '100%' }}
                onChange={(e) => {
                  setPatchUser((prev) => {
                    const newUser = { ...prev };
                    newUser.dateOfBirth = e.target.value;

                    return newUser;
                  });
                }}
                value={patchUser.dateOfBirth}
              ></TextField>
            </Grid>
            <Grid item xs={8} sx={{ marginBottom: '1rem' }}>
              <TextField
                label="createDate"
                variant="outlined"
                sx={{ width: '100%' }}
                disabled
                value={patchUser.createDate}
              ></TextField>
            </Grid>
          </Grid>
        </Grid>
        {/* ユーザー検索・更新ここまで */}

        {/* ユーザー新規作成ここから */}
        <Grid
          item
          xs={5}
          sx={{
            backgroundColor: 'azure',
            minHeight: '450px',
          }}
        >
          <Typography variant="h5" textAlign={'left'}>
            ユーザー新規作成
          </Typography>
          <Grid container sx={{ paddingTop: '1rem' }} justifyContent={'left'}>
            <Grid item xs={8} sx={{ marginBottom: '1rem' }}>
              <TextField
                label="FirstName"
                variant="outlined"
                sx={{ width: '100%' }}
                onChange={(e) => {
                  setCreateUser((prev) => {
                    const newUser = { ...prev };
                    newUser.firstName = e.target.value;

                    return newUser;
                  });
                }}
                value={createUser.firstName}
              ></TextField>
            </Grid>
            <Grid item xs={4} sx={{ marginBottom: '1rem' }}>
              <Button
                variant="outlined"
                sx={{ height: '100%' }}
                onClick={async () => {
                  await handleOnClickCreate();
                }}
              >
                Create
              </Button>
            </Grid>
            <Grid item xs={8} sx={{ marginBottom: '1rem' }}>
              <TextField
                label="LastName"
                variant="outlined"
                sx={{ width: '100%' }}
                onChange={(e) => {
                  setCreateUser((prev) => {
                    const newUser = { ...prev };
                    newUser.lastName = e.target.value;

                    return newUser;
                  });
                }}
                value={createUser.lastName}
              ></TextField>
            </Grid>
            <Grid item xs={8} sx={{ marginBottom: '1rem' }}>
              <TextField
                label="Email"
                variant="outlined"
                sx={{ width: '100%' }}
                onChange={(e) => {
                  setCreateUser((prev) => {
                    const newUser = { ...prev };
                    newUser.email = e.target.value;

                    return newUser;
                  });
                }}
                value={createUser.email}
              ></TextField>
            </Grid>
            <Grid item xs={8} sx={{ marginBottom: '1rem' }}>
              <TextField
                label="DateOfBirth"
                variant="outlined"
                sx={{ width: '100%' }}
                onChange={(e) => {
                  setCreateUser((prev) => {
                    const newUser = { ...prev };
                    newUser.dateOfBirth = e.target.value;

                    return newUser;
                  });
                }}
                value={createUser.dateOfBirth}
              ></TextField>
            </Grid>
          </Grid>
        </Grid>
        {/* ユーザー新規作成ここまで */}

        {/* ペット一覧ここから */}
        <Grid
          item
          xs={5}
          sx={{
            backgroundColor: 'azure',
            marginTop: '2rem',
          }}
        >
          <Typography variant="h5" textAlign={'left'}>
            ペット一覧
          </Typography>
          <Grid container sx={{ paddingTop: '1rem' }} justifyContent={'left'}>
            <Grid item xs={8} sx={{ marginBottom: '1rem' }}>
              <TextField
                label="Tags"
                variant="outlined"
                sx={{ width: '100%' }}
                onChange={(e) => {
                  setTags(e.target.value.split(' '));
                }}
                value={tags.join(' ')}
              ></TextField>
            </Grid>
            <Grid item xs={4} sx={{ marginBottom: '1rem' }}>
              <Button
                variant="outlined"
                sx={{ height: '100%' }}
                onClick={async () => {
                  await handleOnClickPetSearch();
                }}
              >
                Search
              </Button>
            </Grid>
          </Grid>
          <Grid container sx={{ paddingTop: '1rem' }} justifyContent={'left'}>
            <Grid item xs={8} sx={{ marginBottom: '1rem' }}>
              <TextField
                label="Limit"
                variant="outlined"
                sx={{ width: '100%' }}
                onChange={(e) => {
                  const limit = parseInt(e.target.value, 10);
                  if (isNaN(limit)) {
                    setLimit(null);
                  } else {
                    setLimit(limit);
                  }
                }}
                value={limit ?? ''}
              ></TextField>
            </Grid>
          </Grid>
          <Grid container sx={{ paddingTop: '1rem' }} justifyContent={'left'}>
            <Grid item xs={12} sx={{ marginBottom: '1rem' }}>
              <TableContainer>
                <Table aria-label="simple table">
                  <TableHead>
                    <TableRow>
                      <TableCell align="left">ID</TableCell>
                      <TableCell align="left">Name</TableCell>
                      <TableCell align="left">Tag</TableCell>
                    </TableRow>
                  </TableHead>
                  <TableBody>
                    {pets.map((pet) => (
                      <TableRow
                        key={pet.id}
                        sx={{
                          '&:last-child td, &:last-child th': { border: 0 },
                        }}
                      >
                        <TableCell align="left">{pet.id}</TableCell>
                        <TableCell align="left">{pet.name}</TableCell>
                        <TableCell align="left">{pet.tag}</TableCell>
                      </TableRow>
                    ))}
                  </TableBody>
                </Table>
              </TableContainer>
            </Grid>
          </Grid>
        </Grid>
      </Grid>
    </Box>
  );
};

export default App;
