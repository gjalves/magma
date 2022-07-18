#include <stdio.h>
#include <dlfcn.h>
#include <err.h>
#include <stdlib.h>
#include <dirent.h>
#include <string.h>

#define PLUGIN_BASEDIR "/usr/share/magma/plugins"

int (*trigger_ue_attach)(const char *imsi);
int (*trigger_ue_detach)(const char *imsi);
static void *dl_mme_trigger = NULL;

static int dummy1(const char *imsi)
{
    return 0;
}

static void plugins_init(void)
{
    DIR *d_magma_plugins;
    struct dirent *d_magma_plugins_entry;

    // Initialize hooks
    trigger_ue_attach = (int (*)(const char *))dummy1;
    trigger_ue_detach = (int (*)(const char *))dummy1;

    if((d_magma_plugins = opendir(PLUGIN_BASEDIR)) == NULL) return;

    while((d_magma_plugins_entry = readdir(d_magma_plugins))) {
        char buf[BUFSIZ];

        if(strlen(d_magma_plugins_entry->d_name) < 4) continue;
        sprintf(buf, "%s/%s", PLUGIN_BASEDIR, d_magma_plugins_entry->d_name);
        if((dl_mme_trigger = dlopen(buf, RTLD_NOW)) == NULL) continue;

        trigger_ue_attach = dlsym(dl_mme_trigger, "trigger_ue_attach");
        trigger_ue_detach = dlsym(dl_mme_trigger, "trigger_ue_detach");
    }

    closedir(d_magma_plugins);
}

static void plugins_fini(void)
{
    if(dl_mme_trigger)
        dlclose(dl_mme_trigger);
}

static void plugins_init(void) __attribute__ ((constructor));
static void plugins_fini(void) __attribute__ ((destructor));
