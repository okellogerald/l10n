Explains the steps of extracting [MethodGroup]s from json groups

Let's assume this example
In `profile/en.json` file:

```json
[
  {
    "_id": "Profile",
    "_name": "profile",
    "name": "name",
    "@name": {
      "description": "A profile name"
    },
    "age": "age"
  },
  {
    "_id": "Account",
    "_name": "account",
    "name": "Account Name",
    "number": "Account Number"
  }
]
```

In `profile/sw.json` file

```json
[
  {
    "name": "Jina",
    "age": "Umri"
  },
  {
    "name": "Jina la Akaunti",
    "number": "Namba ya Akaunti"
  }
]
```

### step 1: Creating Partial Groups from the main-locale file
We'll have groups like

```json
[
  {
    "id": "Profile",
    "name": "profile",
    "methods": [
      {
        "name": "name",
        "placeholders": [],
        "description": "A profile name"
      },
      {
        "name": "age",
        "placeholders": [],
        "description": ""
      }
    ]
  },
  {
    "id": "Account",
    "name": "account",
    "methods": [
      {
        "name": "name",
        "placeholders": [],
        "description": ""
      },
      {
        "name": "number",
        "placeholders": [],
        "description": ""
      }
    ]
  }
]
```

### step 2: Converting locale lists in step 0 to locale map

The final map will be like

```json
{
  "en": {
    "Profile": {
      "name": "Name",
      "age": "Age"
    },
    "Account": {
      "name": "Account Name",
      "number": "Account Number"
    }
  },
  "sw": {
    "Profile": {
      "name": "Jina",
      "age": "Umri"
    },
    "Account": {
      "name": "Jina la akaunti",
      "number": "Namba ya akaunti"
    }
  }
}
```

### step 3: creating complete groups from the partial groups created in step 1 and the translations in step 2

So we loop over every partial groups and add a list of translations for every method

The end result will be like

```json
[
  {
    "id": "Profile",
    "name": "profile",
    "methods": [
      {
        "name": "name",
        "placeholders": [],
        "description": "A profile name",
        "translations": ["Name", "Jina"]
      },
      {
        "name": "age",
        "placeholders": [],
        "description": "",
        "translations": ["Age", "Umri"]
      }
    ]
  },
  {
    "id": "Account",
    "name": "account",
    "methods": [
      {
        "name": "name",
        "placeholders": [],
        "description": "",
        "translations": ["Account Name", "Jina la akaunti"]
      },
      {
        "name": "number",
        "placeholders": [],
        "description": "",
        "translations": ["Account Number", "Namba ya akaunti"]
      }
    ]
  }
]
```
