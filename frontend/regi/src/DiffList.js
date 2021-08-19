
import * as React from "react";
import {useEffect, useState} from "react"
import { List, Datagrid, TextField, NumberField, DateField, 
    TopToolbar,
    Button,} from 'react-admin';
import IconEvent from '@material-ui/icons/Event';

async function postEnable() {
    const  response = await fetch("/api/enable", {
        method: "POST",
    })
    console.log(response)
    return response
}
async function postDisable() {
    const  response = await fetch("/api/disable", {
        method: "POST",
    })
    console.log(response)
    return response
}
async function getEnabled() {
    const  response = await fetch("/api/enabled")
    return response.json()
}

const ListActions = (props) => {
    const [enabled, setEnabled] = useState(false)
    const updateEnabled = ()=>{
        getEnabled().then(data => {
            console.log(data)
            setEnabled(data["enabled"])
        })
    }

    useEffect(updateEnabled, [])

    return <TopToolbar>
        <Button
            onClick={() => { 
                if (enabled) postDisable().then(_=>updateEnabled()).catch(err=>{console.log(err);updateEnabled()})
                else postEnable().then(_=>updateEnabled()).catch(err=>{console.log(err);updateEnabled()})
            }}
            label={enabled ? "disable" : "enable"}
        /> : 
    </TopToolbar>
}
export const DiffList = props => (
    <div>
    <List {...props} actions={<ListActions/>}>
        <Datagrid >
            <TextField source="id" />
            <TextField source="path" />
            <TextField source="request" />
            <NumberField source="status" />
            <TextField source="body" />
            <TextField source="body_golden" />
            <TextField source="status_golden" />
            <DateField source="created_at" />
        </Datagrid>
    </List>
    
    </div>
);