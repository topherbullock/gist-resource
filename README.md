# gist-resource
A Concourse resource for github gists


# Usage

Add the following to your pipeline's `resource_types`

```
- name: gist
  type: docker-image
  source:
    repository: topherbullock/gist-resource
```


##Source Configuration

- `id`: Required. The github gist ID

- `access_token`: Optional. the Github Access Token to use, with `gist` permissions


##Behaviour

### `check`: Check for new gist revisions

### `in`: Get all the files for a revision of a gist 


