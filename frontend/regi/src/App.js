
import * as React from "react";
import jsonServerProvider from 'ra-data-json-server';
import { Admin, Resource, ListGuesser } from 'react-admin';
import {DiffList} from "./DiffList"
const dataProvider = jsonServerProvider('api');
const App = () => (
      <Admin dataProvider={dataProvider}>
          <Resource name="diffs" list={DiffList} />
      </Admin>
  );
  
export default App;
